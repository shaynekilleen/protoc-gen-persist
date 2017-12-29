// This file is generated by protoc-gen-persist
// Source File: examples/spanner/bob_example/bobs.proto
// DO NOT EDIT !
package bob_example

import (
	io "io"

	spanner "cloud.google.com/go/spanner"
	mytime "github.com/tcncloud/protoc-gen-persist/examples/mytime"
	persist_lib "github.com/tcncloud/protoc-gen-persist/examples/spanner/bob_example/persist_lib"
	context "golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

// don't eat our spanner import, or complain
var _ = spanner.NewClient

type BobsImpl struct {
	PERSIST   *persist_lib.BobsMethodReceiver
	FORWARDED RestOfBobsHandlers
}
type RestOfBobsHandlers interface {
}
type BobsImplBuilder struct {
	err           error
	rest          RestOfBobsHandlers
	queryHandlers *persist_lib.BobsQueryHandlers
	i             *BobsImpl
	db            *spanner.Client
}

func NewBobsBuilder() *BobsImplBuilder {
	return &BobsImplBuilder{i: &BobsImpl{}}
}
func (b *BobsImplBuilder) WithRestOfGrpcHandlers(r RestOfBobsHandlers) *BobsImplBuilder {
	b.rest = r
	return b
}
func (b *BobsImplBuilder) WithPersistQueryHandlers(p *persist_lib.BobsQueryHandlers) *BobsImplBuilder {
	b.queryHandlers = p
	return b
}
func (b *BobsImplBuilder) WithDefaultQueryHandlers() *BobsImplBuilder {
	accessor := persist_lib.NewSpannerClientGetter(b.db)
	queryHandlers := &persist_lib.BobsQueryHandlers{
		DeleteBobsHandler:         persist_lib.DefaultDeleteBobsHandler(accessor),
		PutBobsHandler:            persist_lib.DefaultPutBobsHandler(accessor),
		GetBobsHandler:            persist_lib.DefaultGetBobsHandler(accessor),
		GetPeopleFromNamesHandler: persist_lib.DefaultGetPeopleFromNamesHandler(accessor),
	}
	b.queryHandlers = queryHandlers
	return b
}
func (b *BobsImplBuilder) WithSpannerClient(c *spanner.Client) *BobsImplBuilder {
	b.db = c
	return b
}
func (b *BobsImplBuilder) WithSpannerURI(ctx context.Context, uri string) *BobsImplBuilder {
	cli, err := spanner.NewClient(ctx, uri)
	b.err = err
	b.db = cli
	return b
}
func (b *BobsImplBuilder) Build() (*BobsImpl, error) {
	if b.err != nil {
		return nil, b.err
	}
	b.i.PERSIST = &persist_lib.BobsMethodReceiver{Handlers: *b.queryHandlers}
	b.i.FORWARDED = b.rest
	return b.i, nil
}

func (s *BobsImpl) DeleteBobs(ctx context.Context, req *Bob) (*Empty, error) {
	var err error
	_ = err
	params := &persist_lib.BobForBobs{}
	// set 'Bob.start_time' in params
	if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "could not convert type to persist_lib type: %v, err", err)
	}
	var res = Empty{}
	var iterErr error
	_ = iterErr
	_ = res
	err = s.PERSIST.DeleteBobs(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
	})
	if err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "error in closure: %v", err)
	}
	return &res, nil
}

func (s *BobsImpl) PutBobs(stream Bobs_PutBobsServer) error {
	var err error
	_ = err
	feed, stop := s.PERSIST.PutBobs(stream.Context())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, "error recieving input: %v", err)
		}
		params := &persist_lib.BobForBobs{}
		// set 'Bob.id' in params
		params.Id = req.Id
		// set 'Bob.name' in params
		params.Name = req.Name
		// set 'Bob.start_time' in params
		if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
			return gstatus.Errorf(codes.Unknown, "could not convert type to persist_lib type: %v, err", err)
		}
		feed(params)
	}
	row, err := stop()
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error receiving result row: %v", err)
	}
	res := NumRows{}
	if row != nil {
		var Count int64
		{
			local := &spanner.NullInt64{}
			if err := row.ColumnByName("count", local); err != nil {
				return gstatus.Errorf(codes.Unknown, "couldnt scan out message, err: %v", err)
			}
			if local.Valid {
				Count = local.Int64
			}
			res.Count = Count
		}
	}
	if err := stream.SendAndClose(&res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error sending response: %v", err)
	}
	return nil
}

func (s *BobsImpl) GetBobs(req *Empty, stream Bobs_GetBobsServer) error {
	var err error
	_ = err
	params := &persist_lib.EmptyForBobs{}
	var iterErr error
	err = s.PERSIST.GetBobs(stream.Context(), params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		res := Bob{}
		var Id int64
		{
			local := &spanner.NullInt64{}
			if err := row.ColumnByName("id", local); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
			}
			if local.Valid {
				Id = local.Int64
			}
			res.Id = Id
		}
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
		}
		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		{
			local := &spanner.NullString{}
			if err := row.ColumnByName("name", local); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
			}
			if local.Valid {
				Name = local.StringVal
			}
			res.Name = Name
		}
		if err := stream.Send(&res); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
		}
	})
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error during iteration: %v", err)
	} else if iterErr != nil {
		return iterErr
	}
	return nil
}

func (s *BobsImpl) GetPeopleFromNames(req *Names, stream Bobs_GetPeopleFromNamesServer) error {
	var err error
	_ = err
	params := &persist_lib.NamesForBobs{}
	// set 'Names.names' in params
	params.Names = req.Names
	var iterErr error
	err = s.PERSIST.GetPeopleFromNames(stream.Context(), params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		res := Bob{}
		var Id int64
		{
			local := &spanner.NullInt64{}
			if err := row.ColumnByName("id", local); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
			}
			if local.Valid {
				Id = local.Int64
			}
			res.Id = Id
		}
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
		}
		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		{
			local := &spanner.NullString{}
			if err := row.ColumnByName("name", local); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
			}
			if local.Valid {
				Name = local.StringVal
			}
			res.Name = Name
		}
		if err := stream.Send(&res); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "couldnt scan out message err: %v", err)
		}
	})
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error during iteration: %v", err)
	} else if iterErr != nil {
		return iterErr
	}
	return nil
}
