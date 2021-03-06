package generator

import (
	"fmt"
)

type PersistStringer struct{}

// TYPECHANGE
func (per *PersistStringer) MessageInputDeclaration(method *Method) string {
	printer := &Printer{}
	printer.P("type %s struct{\n", NewPLInputName(method))

	getPersistLibTypeName := GetSqlPersistLibTypeName
	if method.IsSpanner() {
		getPersistLibTypeName = GetSpannerPersistLibTypeName
	}

	inputTypeDescs := method.GetTypeDescArrayForStruct(method.GetInputTypeStruct())
	for _, qf := range inputTypeDescs {
		typeName := getPersistLibTypeName(qf)
		printer.P("%s %s\n", qf.Name, typeName)
	}
	printer.P("}\n")
	printer.P("// this could be used in a query, so generate the getters/setters\n")
	for _, qf := range inputTypeDescs {
		typeName := getPersistLibTypeName(qf)
		plInputName := NewPLInputName(method)
		printer.P(
			"func(p *%s) Get%s() %s{ return p.%s }\n",
			plInputName, qf.Name, typeName, qf.Name,
		)
		printer.P(
			"func(p *%s) Set%s(param %s) { p.%s = param }\n",
			plInputName, qf.Name, typeName, qf.Name,
		)
	}
	return printer.String()
}

// a cache for the method types we have already written
func (per *PersistStringer) PersistImplBuilder(service *Service, alreadyWrote map[string]bool) string {
	var dbType string
	var backend string
	if service.IsSpanner() {
		dbType = "spanner.Client"
		backend = "Spanner"
	} else {
		dbType = "sql.DB"
		backend = "Sql"
	}
	sName := service.GetName()
	printer := &Printer{}
	printer.Q(
		"type ", sName, "Impl struct{\n",
		"PERSIST *persist_lib.", NewPersistHelperName(service), "\n",
		"FORWARDED RestOf", sName, "Handlers\n",
		"HOOKS ", sName, "Hooks\n",
		"MAPPINGS ", sName, "TypeMapping\n",
		"}\n",
	)
	printer.P("type RestOf%sHandlers interface{\n", service.GetName())
	for _, m := range *service.Methods {
		spannerBi := m.Service.IsSpanner() && m.IsBidiStreaming()
		if m.GetMethodOption() == nil || spannerBi {
			if m.IsUnary() {
				printer.P(
					"%s(ctx context.Context, req *%s) (*%s, error)\n",
					m.GetName(),
					m.GetInputType(),
					m.GetOutputType(),
				)
			} else if m.IsServerStreaming() {
				printer.P(
					"%s(req *%s, stream %s) error\n",
					m.GetName(),
					m.GetInputType(),
					NewStreamType(m),
				)
			} else {
				printer.P(
					"%s(stream %s) error\n",
					m.GetName(),
					NewStreamType(m),
				)
			}
		}
	}
	printer.P("}\n")
	WriteBuilderTypeMappingsInterface(printer, service)
	WriteTypeMappingsContractInterfaces(printer, service, alreadyWrote)
	WriteBuilderHookInterfaceAndFunc(printer, service)

	printer.Q(
		"type ", sName, "ImplBuilder struct {\n",
		"err error\n ",
		"rest RestOf", sName, "Handlers\n",
		"queryHandlers *persist_lib.", sName, "QueryHandlers\n",
		"i *", sName, "Impl\n",
		"db *", dbType, "\n",
		"hooks ", sName, "Hooks\n",
		"mappings ", sName, "TypeMapping\n",
		"}\n",
		"func New", sName, "Builder() *", sName, "ImplBuilder {\n",
		"return &", sName, "ImplBuilder{i: &", sName, "Impl{}}\n",
		"}\n",
	)

	WriteBuilderHooksAcceptingFunc(printer, service)
	WriteBuilderTypeMappingsAcceptingFunc(printer, service)

	printer.PA([]string{
		"func (b *%sImplBuilder) WithRestOfGrpcHandlers(r RestOf%sHandlers) *%sImplBuilder {\n",
		"b.rest = r\n return b\n}\n",
	},
		service.GetName(),
		service.GetName(),
		service.GetName(),
	)
	printer.PA([]string{
		"func (b *%sImplBuilder) WithPersistQueryHandlers(p *persist_lib.%sQueryHandlers)",
		"*%sImplBuilder {\n",
		"b.queryHandlers = p\n return b\n}\n",
	},
		service.GetName(),
		service.GetName(),
		service.GetName(),
	)

	// setup default query functions
	printer.PA([]string{
		"func (b *%sImplBuilder) WithDefaultQueryHandlers() *%sImplBuilder {\n",
		"accessor := persist_lib.New%sClientGetter(&b.db)\n",
		"queryHandlers := &persist_lib.%sQueryHandlers{\n",
	},
		service.GetName(), service.GetName(),
		backend,
		service.GetName(),
	)
	for _, m := range *service.Methods {
		if m.GetMethodOption() == nil || (m.Service.IsSpanner() && m.IsBidiStreaming()) {
			continue
		}
		printer.P(
			"%s: persist_lib.Default%s(accessor),\n",
			NewPersistHandlerName(m),
			NewPersistHandlerName(m),
		)
	}
	printer.P("}\n b.queryHandlers = queryHandlers\n return b\n}\n")
	// fill in holes with defaults
	printer.PA([]string{
		"func (b *%sImplBuilder) WithNilAsDefaultQueryHandlers(p *persist_lib.%sQueryHandlers)",
		"*%sImplBuilder {\n",
		"accessor := persist_lib.New%sClientGetter(&b.db)\n",
	},
		service.GetName(), service.GetName(),
		service.GetName(),
		backend,
	)
	for _, m := range *service.Methods {
		if m.GetMethodOption() == nil || (m.Service.IsSpanner() && m.IsBidiStreaming()) {
			continue
		}
		phn := NewPersistHandlerName(m)
		printer.P(
			"if p.%s == nil {\np.%s = persist_lib.Default%s(accessor)\n}\n",
			phn, phn, phn,
		)
	}
	printer.P("b.queryHandlers = p\n return b\n}\n")

	// provide the builder with a client
	printer.PA([]string{
		"func (b *%sImplBuilder) With%sClient(c *%s) *%sImplBuilder {\n",
		"b.db = c\n return b\n}\n",
	},
		service.GetName(), backend, dbType, service.GetName(),
	)

	if service.IsSpanner() {
		printer.PA([]string{
			"func (b *%sImplBuilder) WithSpannerURI(ctx context.Context, uri string) *%sImplBuilder {\n",
			"cli, err := spanner.NewClient(ctx, uri)\n b.err = err\n b.db = cli\n return b\n}\n",
		},
			service.GetName(), service.GetName(),
		)
	} else {
		printer.PA([]string{
			"func (b *%sImplBuilder) WithNewSqlDb(driverName, dataSourceName string) *%sImplBuilder {\n",
			"db, err := sql.Open(driverName, dataSourceName)\n",
			"b.err = err\n",
			"if b.err == nil {\n",
			"\tb.db = db\n",
			"}\n",
			"return b\n}\n",
		},
			service.GetName(), service.GetName(),
		)
	}
	// Build method, returns impl, err
	printer.Q(
		"func (b *", sName, "ImplBuilder) Build() (*", sName, "Impl, error) {\n",
		"if b.err != nil {\n return nil, b.err\n",
		"}\n",
		"b.i.PERSIST = &persist_lib.", NewPersistHelperName(service), "{Handlers: *b.queryHandlers}\n",
		"b.i.FORWARDED = b.rest\n",
		"b.i.HOOKS = b.hooks\n",
		"b.i.MAPPINGS = b.mappings\n",
		"return b.i, nil\n",
		"}\n",
	)
	// MustBuild method, returns impl.  Can panic.
	printer.PA([]string{
		"func (b *%sImplBuilder) MustBuild() *%sImpl {\n",
		"s, err := b.Build()\n",
		"if err != nil {\n panic(\"error in builder: \" + err.Error())\n}\n",
		"return s\n}\n",
	},
		service.GetName(), service.GetName(),
	)
	return printer.String()
}

func WriteBuilderHookInterfaceAndFunc(p *Printer, s *Service) {
	p.Q("type ", s.GetName(), "Hooks interface{\n")
	for _, m := range *s.Methods {
		opt := m.GetMethodOption()
		if opt == nil {
			continue
		}
		if opt.GetBefore() {
			sliceStarOrStar := "*"
			if m.IsServerStreaming() {
				sliceStarOrStar = "[]*"
			}

			p.Q("\t", m.GetBeforeHookName(), "(*", m.GetInputType(), ") (", sliceStarOrStar, m.GetOutputType(), ", error)\n")
		}
		if opt.GetAfter() {
			p.Q("\t", m.GetAfterHookName(), "(*", m.GetInputType(), ", *", m.GetOutputType(), ") error\n")
		}
	}
	p.Q("}\n")
}
func WriteBuilderHooksAcceptingFunc(p *Printer, serv *Service) {
	s := serv.GetName()
	p.Q(
		"func(b *", s, "ImplBuilder) WithHooks(hs ", s, "Hooks) *", s, "ImplBuilder {\n",
		"b.hooks = hs\n",
		"return b\n",
		"}\n",
	)
}

func WriteBuilderTypeMappingsAcceptingFunc(p *Printer, serv *Service) {
	s := serv.GetName()
	p.Q("func(b *", s, "ImplBuilder) WithTypeMapping(ts ", s, "TypeMapping) *", s, "ImplBuilder {\n")
	p.Q("\tb.mappings = ts\n")
	p.Q("\treturn b\n")
	p.Q("}\n")
}

func WriteBuilderTypeMappingsInterface(p *Printer, s *Service) {
	sName := s.GetName()
	// TODO google's WKT protobufs probably don't need the package prefix
	p.Q("type ", sName, "TypeMapping interface{\n")
	tms := s.GetServiceOption().GetTypes()
	for _, tm := range tms {
		// TODO implement these interfaces
		_, titled := getGoNamesForTypeMapping(tm, s.File)
		// p.Q(titled, "() ", sName, titled, "MappingImpl\n")
		p.Q(titled, "() ", titled, "MappingImpl\n")
	}
	p.Q("}\n")

}
func WriteScanValuerInterface(p *Printer, s *Service) {
	if s.IsSQL() {
		p.Q("type ScanValuer interface {\n")
		p.Q("\tsql.Scanner\n")
		p.Q("\tdriver.Valuer\n")
		p.Q("}\n")
	} else if s.IsSpanner() {
		p.Q("type ScanValuer interface {\n")
		p.Q("\tSpannerScan(src *spanner.GenericColumnValue) error\n")
		p.Q("\tSpannerValue() (interface{}, error)\n")
		p.Q("}\n")
	}
}
func WriteTypeMappingsContractInterfaces(p *Printer, s *Service, alreadyWrote map[string]bool) {
	for _, tm := range s.GetServiceOption().GetTypes() {
		name, titled := getGoNamesForTypeMapping(tm, s.File)
		if alreadyWrote[titled] {
			continue
		}
		_, maybeStar := needsExtraStar(tm)
		p.Q("type ", titled, "MappingImpl interface{\n")
		p.Q("ToProto(*", maybeStar, name, ") error\n")
		p.Q("Empty() ", titled, "MappingImpl\n")
		if s.IsSQL() {
			p.Q("ToSql(", maybeStar, name, ") sql.Scanner\n")
			p.Q("sql.Scanner\n")
			p.Q("driver.Valuer\n")
		} else if s.IsSpanner() {
			p.Q("ToSpanner(", maybeStar, name, ") ", titled, "MappingImpl\n")
			p.Q("SpannerScan(src *spanner.GenericColumnValue) error\n")
			p.Q("SpannerValue() (interface{}, error)\n")
		}
		p.Q("}\n")
		alreadyWrote[titled] = true
	}
}

func (per *PersistStringer) HandlersStructDeclaration(service *Service) string {
	printer := &Printer{}

	// contains our query handlers struct, and is reciever of our methods
	printer.P(
		"type %s struct{\nHandlers %sQueryHandlers}\n",
		NewPersistHelperName(service),
		service.GetName(),
	)
	// actually runs the queries
	printer.P("type %sQueryHandlers struct {\n", service.GetName())
	for _, method := range *service.Methods {
		if method.GetMethodOption() == nil {
			continue
		}
		var rowType string
		if method.IsSpanner() {
			rowType = "*spanner.Row"
		} else {
			rowType = "Scanable"
		}
		if method.IsClientStreaming() {
			printer.P(
				"%s func(context.Context)(func(*%s)error, func() (%s, error), error)\n",
				NewPersistHandlerName(method),
				NewPLInputName(method),
				rowType,
			)
		} else if method.IsBidiStreaming() {
			printer.P(
				"%s func(context.Context) (func(*%s) (%s, error), func() error)\n",
				NewPersistHandlerName(method),
				NewPLInputName(method),
				rowType,
			)
		} else {
			printer.P(
				"%s func(context.Context, *%s, func(%s)) error\n",
				NewPersistHandlerName(method),
				NewPLInputName(method),
				rowType,
			)
		}
	}
	printer.P("}\n")
	return printer.String()
}

func (per *PersistStringer) HelperFunctionImpl(service *Service) string {
	printer := &Printer{}
	for _, method := range *service.Methods {
		if method.GetMethodOption() == nil {
			continue // we do not have any persist options
		}
		var rowType string
		if method.IsSpanner() {
			rowType = "*spanner.Row"
		} else {
			rowType = "Scanable"
		}

		if method.IsClientStreaming() {
			printer.PA([]string{
				"// given a context, returns two functions.  (feed, stop)\n",
				"// feed will be called once for every row recieved by the handler\n",
				"// stop will be called when the client is done streaming. it expects\n",
				"//a  row to be returned, or nil.\n",
				"func (p *%s) %s(ctx context.Context)(func(*%s) error, func() (%s, error), error) {\n",
				"return p.Handlers.%s(ctx)\n}\n",
			},
				NewPersistHelperName(service),
				method.GetName(),
				NewPLInputName(method),
				rowType,
				NewPersistHandlerName(method),
			)
		} else if method.IsBidiStreaming() {
			printer.PA([]string{
				"// returns two functions (feed, stop)\n",
				"// feed needs to be called for every row received. It will run the query\n",
				"// and return the result + error",
				"// stop needs to be called to signal the transaction has finished\n",
				"func (p *%s) %s(ctx context.Context)(func(*%s) (%s, error), func() error) {\n",
				"return p.Handlers.%s(ctx)\n}\n",
			},
				NewPersistHelperName(service),
				method.GetName(),
				NewPLInputName(method),
				rowType,
				NewPersistHandlerName(method),
			)
		} else {
			printer.PA([]string{
				"// next must be called on each result row\n",
				"func(p *%s) %s(ctx context.Context, params *%s, next func(%s)) error {\n",
				"return p.Handlers.%s(ctx, params, next)\n}\n",
			},
				NewPersistHelperName(service),
				method.GetName(),
				NewPLInputName(method),
				rowType,
				NewPersistHandlerName(method),
			)
		}
	}
	return printer.String()
}

func (per *PersistStringer) QueryInterfaceDefinition(method *Method) string {
	if method.GetMethodOption() == nil {
		return ""
	}
	printer := &Printer{}
	printer.P(
		"type %sParams interface{\n",
		NewPLQueryMethodName(method),
	)
	getPersistLibTypeName := GetSpannerPersistLibTypeName
	if method.IsSQL() {
		getPersistLibTypeName = GetSqlPersistLibTypeName
	}

	for _, t := range method.GetTypeDescForQueryFields() {
		interfaceType := getPersistLibTypeName(t)
		printer.P("Get%s() %s\n", t.Name, interfaceType)
	}
	printer.P("}\n")
	return printer.String()
}

// TYPECHANGE
func (per *PersistStringer) SqlQueryFunction(method *Method) string {
	opts := method.GetMethodOption()
	if opts == nil {
		return ""
	}
	// Join query with space
	query := func() (out string) {
		for _, q := range opts.GetQuery() {
			out += q + " "
		}
		return
	}()

	args := opts.GetArguments()
	tds := method.GetTypeDescForFieldsInStructSnakeCase(method.GetInputTypeStruct())

	var argParams string
	for _, a := range args {
		argParams += fmt.Sprintf("req.Get%s(),\n", tds[a].Name)
	}
	// if we are an empty result, then perform an exec, not a query
	lenOfResult := len(method.GetTypeDescArrayForStruct(method.GetOutputTypeStruct()))

	printer := &Printer{}
	queryMethodName := NewPLQueryMethodName(method)
	printer.P("func %s(tx Runable, req %sParams) *Result {", queryMethodName, queryMethodName)
	if lenOfResult == 0 || method.IsClientStreaming() { // use an exec
		printer.PA([]string{
			"res, err := tx.Exec(\n\"%s\",\n%s)\n",
			"if err != nil {\n return newResultFromErr(err)\n}\n",
			"return newResultFromSqlResult(res)\n",
		},
			query, argParams,
		)
	} else if method.IsServerStreaming() {
		printer.PA([]string{
			"res, err := tx.Query(\n\"%s\",\n%s)\n",
			"if err != nil {\n return newResultFromErr(err)\n}\n",
			"return newResultFromRows(res)",
		},
			query, argParams,
		)
	} else {
		printer.PA([]string{
			"row := tx.QueryRow(\n\"%s\",\n%s)\n",
			"return newResultFromRow(row)\n",
		},
			query, argParams,
		)
	}
	printer.P("\n}\n")

	return printer.String()
}
func (per *PersistStringer) SpannerQueryFunction(method *Method) string {
	// we do not have a persist query
	if method.GetMethodOption() == nil {
		return ""
	}
	printer := &Printer{}
	if method.IsSelect() {
		printer.P(
			"func %s(req %sParams) spanner.Statement {\nreturn %s\n}\n",
			NewPLQueryMethodName(method),
			NewPLQueryMethodName(method),
			method.Query,
		)
	} else {
		printer.P(
			"func %s(req %sParams) *spanner.Mutation {\nreturn %s\n}\n",
			NewPLQueryMethodName(method),
			NewPLQueryMethodName(method),
			method.Query,
		)
	}
	return printer.String()
}
func (per *PersistStringer) DefaultFunctionsImpl(service *Service) string {
	printer := &Printer{}
	for _, method := range *service.Methods {
		if method.GetMethodOption() == nil {
			continue
		}
		if method.IsSQL() {
			printer.P("%s", per.DefaultSqlFunctionsImpl(method))
		} else if method.IsSpanner() {
			printer.P("%s", per.DefaultSpannerFunctionsImpl(method))
		}
	}
	return printer.String()
}
func (per *PersistStringer) DefaultSqlFunctionsImpl(method *Method) string {
	printer := &Printer{}
	lenOfOutFields := len(method.GetTypeDescArrayForStruct(method.GetOutputTypeStruct()))
	if method.IsClientStreaming() { // use exec
		printer.PA([]string{
			"func Default%sHandler(accessor SqlClientGetter) func(context.Context) ",
			"(func(*%s) error, func() (Scanable, error), error) {\n",
			"return func(ctx context.Context) (func(*%s) error, func() (Scanable, error), error) {\n",
			"sqlDb, err := accessor()\n",
			"if err != nil {\n return nil, nil, err\n}\n",
			"tx, err := sqlDb.Begin()\n",
			"if err != nil {\n return nil, nil, err\n}\n",
			"feed := func(req *%s) error {\n",
			"if res := %s(tx, req); res.Err() != nil {\n",
			"if err := tx.Rollback(); err != nil {\n return fmt.Errorf(\"%s\", err, res.Err())\n}\n",
			"return res.Err()\n}\n",
			"return nil\n}\n",
			"done := func() (Scanable, error) {\n if err := tx.Commit();err != nil {\n",
			"return nil, err\n}\n return nil, nil\n}\n",
			"return feed, done, nil\n}\n}\n",
		},
			method.GetName(),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLQueryMethodName(method),
			"%v, %v",
		)
	} else if lenOfOutFields == 0 { // use exec
		printer.PA([]string{
			"func Default%sHandler(accessor SqlClientGetter) ",
			"func (context.Context, *%s, func(Scanable)) error {\n",
			"return func(ctx context.Context, req *%s, next func(Scanable)) error {\n",
			"sqlDB, err := accessor()\n if err != nil {\n return err \n}\n",
			"if res := %s(sqlDB, req); res.Err() != nil {\n return err \n}\n",
			"return nil\n}\n}\n",
		},
			method.GetName(),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLQueryMethodName(method),
		)
	} else if method.IsServerStreaming() { // use query
		printer.PA([]string{
			"func Default%sHandler(accessor SqlClientGetter) ",
			"func(context.Context, *%s, func(Scanable)) error {\n",
			"return func(ctx context.Context, req *%s, next func(Scanable)) error {\n",
			"sqlDB, err := accessor()\n if err != nil {\n return err\n}\n",
			"tx, err := sqlDB.Begin()\n",
			"if err != nil {\n return err\n}\n",
			"res := %s(tx, req)\n",
			"err = res.Do(func(row Scanable) error {\n",
			"next(row)\n return nil\n})\n",
			"if err != nil {\n return err \n}\n",
			"if err := tx.Commit(); err != nil { return err \n}\n",
			"return res.Err()\n}\n}\n",
		},
			method.GetName(),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLQueryMethodName(method),
		)
	} else if method.IsUnary() { // use queryRow
		printer.PA([]string{
			"func Default%sHandler(accessor SqlClientGetter)  ",
			"func(context.Context, *%s, func(Scanable)) error {\n",
			"return func(ctx context.Context, req *%s, next func(Scanable)) error {\n",
			"sqlDB, err := accessor()\n if err != nil {\n return err\n}\n",
			"res := %s(sqlDB, req)\n",
			"err = res.Do(func(row Scanable) error {\n",
			"next(row)\nreturn nil})\n",
			"if err != nil {\n return err\n}\n",
			"return nil\n}\n",
			"}\n",
		},
			method.GetName(),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLQueryMethodName(method),
		)
	} else if method.IsBidiStreaming() {
		printer.PA([]string{
			"func Default%sHandler(accessor SqlClientGetter) ",
			"func(context.Context) (func(*%s) (Scanable, error), func() error) {\n",
			"return func(ctx context.Context) (func(*%s) (Scanable, error), func() error) {\n",
			"var feedErr error\n",
			"sqlDb, err := accessor()\n",
			"if err != nil {\n feedErr = err\n}\n",
			"tx, err := sqlDb.Begin()\n",
			"if err != nil {\n feedErr = err\n}\n",
			"feed := func(req *%s) (Scanable, error) {\n",
			"if feedErr != nil{\n return nil, feedErr\n}\n res := %s(tx, req)\n",
			"return res, nil\n}\n",
			"done := func() error {\n if feedErr != nil {\n tx.Rollback()\n} else {\n feedErr = tx.Commit()\n}\n",
			"return feedErr\n}\n",
			"return feed,done\n}\n}\n",
		},
			method.GetName(),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLQueryMethodName(method),
		)
	}
	return printer.String()
}
func (per *PersistStringer) DefaultSpannerFunctionsImpl(method *Method) string {
	printer := &Printer{}
	if method.IsClientStreaming() {
		printer.PA([]string{
			"func Default%sHandler(accessor SpannerClientGetter) func(context.Context) ",
			"(func(*%s) error, func()(*spanner.Row, error), error) {\n",
			"return func(ctx context.Context) (func(*%s) error, func()(*spanner.Row, error), error) {\n",
			"var muts []*spanner.Mutation\n",
			"feed := func(req *%s) error {\nmuts = append(muts, %s(req))\nreturn nil\n}\n",
			"done := func() (*spanner.Row, error) {\n",
			"cli, err := accessor()\nif err != nil {\n return nil, err\n}\n",
			"if _, err := cli.Apply(ctx, muts); err != nil {\nreturn nil, err\n}\n",
			"return nil, nil // we dont have a row, because we are an apply\n",
			"}\n return feed, done, nil\n}\n}\n",
		},
			method.GetName(),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLInputName(method),
			NewPLQueryMethodName(method),
		)
	} else {
		printer.PA([]string{
			"func Default%sHandler(accessor SpannerClientGetter) ",
			"func(context.Context, *%s, func(*spanner.Row)) error {\n",
			"return func(ctx context.Context, req *%s, next func(*spanner.Row)) error {\n",
		},
			method.GetName(),
			NewPLInputName(method),
			NewPLInputName(method),
		)
		printer.P("cli, err := accessor()\n if err != nil {\n return err\n}\n")

		if method.IsSelect() {
			printer.PA([]string{
				"iter := cli.Single().Query(ctx, %s(req))\n",
				"if err := iter.Do(func(r *spanner.Row) error {\n",
				"next(r)\nreturn nil\n}); err != nil {\nreturn err\n}\n",
			},
				NewPLQueryMethodName(method),
			)
		} else {
			printer.PA([]string{
				"if _, err := cli.Apply(ctx, []*spanner.Mutation{%s(req)}); err != nil {\n",
				"return err\n}\n next(nil) // this is an apply, it has no result\n",
			},
				NewPLQueryMethodName(method),
			)
		}
		printer.P("return nil\n}\n}\n")
	}
	return printer.String()
}

func (per *PersistStringer) DeclareSpannerGetter() string {
	printer := &Printer{}
	printer.P("type SpannerClientGetter func() (*spanner.Client, error)\n")
	printer.PA([]string{
		"func NewSpannerClientGetter(cli **spanner.Client) SpannerClientGetter {\n",
		"return func() (*spanner.Client, error) {\n return *cli, nil \n}\n}\n",
	})

	return printer.String()
}

// package level definitions for sql implemented libraries
// SqlClientGetter
// Scanable interface
// Runable interface
// Result struct
func (per *PersistStringer) DeclareSqlPackageDefs() string {
	printer := &Printer{}
	printer.P("type SqlClientGetter func() (*sql.DB, error)\n")
	printer.PA([]string{
		"func NewSqlClientGetter(cli **sql.DB) SqlClientGetter {\n",
		"return func() (*sql.DB, error) {\n return *cli, nil \n}\n}\n",
	})
	printer.P("type Scanable interface{\nScan(dest ...interface{}) error\n}\n")
	printer.PA([]string{
		"type Runable interface{\n",
		"Query(string, ...interface{}) (*sql.Rows, error)\n",
		"QueryRow(string, ...interface{}) *sql.Row\n",
		"Exec(string, ...interface{}) (sql.Result, error)\n}\n",
	})
	printer.PA([]string{
		"type Result struct {\n",
		"result sql.Result\n",
		"row    *sql.Row\n",
		"rows   *sql.Rows\n",
		"err    error\n",
		"}\n",
		"func newResultFromSqlResult(r sql.Result) *Result {\n",
		"return &Result{result: r}\n",
		"}\n",
		"func newResultFromRow(r *sql.Row) *Result {\n",
		"return &Result{row: r}\n",
		"}\n",
		"func newResultFromRows(r *sql.Rows) *Result {\n",
		"return &Result{rows: r}\n",
		"}\n",
		"func newResultFromErr(err error) *Result {\n",
		"return &Result{err: err}\n",
		"}\n",
		"func (r *Result) Do(fun func(Scanable) error) error {\n",
		"if r.err != nil {\n",
		"return r.err\n",
		"}\n",
		"if r.row != nil {\n",
		"if err := fun(r.row); err != nil {\n",
		"return err\n",
		"}\n",
		"}\n",
		"if r.rows != nil {\n",
		"defer r.rows.Close()\n",
		"for r.rows.Next() {\n",
		"if err := fun(r.rows); err != nil {\n",
		"return err\n",
		"}\n",
		"}\n",
		"}\n",
		"return nil\n",
		"}\n",
		"// returns sql.ErrNoRows if it did not scan into dest\n",
		"func (r *Result) Scan(dest ...interface{}) error {\n",
		"if r.result != nil {\n return sql.ErrNoRows\n}",
		"else if r.row != nil {\n return r.row.Scan(dest...)\n}",
		"else if r.rows != nil {\n",
		"err := r.rows.Scan(dest...)\n",
		"if r.rows.Next() {\n r.rows.Close()\n}\n",
		"return err\n",
		"}\n",
		"return sql.ErrNoRows\n",
		"}\n",
		"func (r *Result) Err() error {\n",
		"return r.err\n",
		"}\n",
	})
	return printer.String()
}

func IteratorHelper(m *Method) string {
	var iterType string

	if m.Service.IsSpanner() {
		iterType = "spanner.RowIterator"
	} else if m.Service.IsSQL() {
		iterType = "persist_lib.Result"
	}
	p := &Printer{}
	out := m.GetOutputType()
	sName := m.Service.GetName()
	p.Q("func ", IterProtoName(m), "(ms ", sName, "TypeMapping, iter *", iterType, ", next func(i *", out, ") error) error {\n")
	p.PA([]string{
		"return iter.Do(func(r %s) error {\n",
		"item, err := %s(ms, r)\n",
		"if err != nil {\n",
		"return fmt.Errorf(\"error converting %s row to protobuf message: %s\", err)\n",
		"}\n",
		"return next(item)\n})\n}\n",
	},
		m.backend.RowType(),
		FromScanableFuncName(m),
		m.GetOutputType(), "%s", // so our printer doesnt freak out
	)
	return p.String()
}

// TYPECHANGE
func GetSqlPersistLibTypeName(t TypeDesc) string {
	if t.IsMapped {
		return "interface{}"
	} else if t.IsMessage {
		return "[]byte"
	} else {
		return t.GoName
	}
}

// TYPECHANGE
func GetSpannerPersistLibTypeName(t TypeDesc) string {
	if t.IsMapped {
		return "interface{}"
	} else if t.IsMessage && t.IsRepeated {
		return "[][]byte"
	} else if t.IsMessage {
		return "[]byte"
	} else {
		return t.GoName
	}
}
