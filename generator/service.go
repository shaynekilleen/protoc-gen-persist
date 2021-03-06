// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/sirupsen/logrus"
	"github.com/tcncloud/protoc-gen-persist/persist"
)

type Service struct {
	Desc       *descriptor.ServiceDescriptorProto
	Methods    *Methods
	Package    string // protobuf package
	File       *FileStruct
	AllStructs *StructList
}

func (s *Service) ProcessMethods() error {
	for _, m := range s.Desc.GetMethod() {
		if err := s.Methods.AddMethod(m, s); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Process() error {
	err := s.ProcessMethods()
	if err != nil {
		return fmt.Errorf("%s\n  service: %s", err, s.GetName())
	}
	return nil
}

func (s *Service) GetName() string {
	return s.Desc.GetName()
}

func (s *Service) GetServiceOption() *persist.TypeMapping {
	if s.Desc.Options != nil && proto.HasExtension(s.Desc.Options, persist.E_Mapping) {
		ext, err := proto.GetExtension(s.Desc.Options, persist.E_Mapping)
		if err == nil {
			return ext.(*persist.TypeMapping)
		}
	}
	return nil
}

func (s *Service) GetServiceType() *persist.PersistenceOptions {
	if s.Desc.Options != nil && proto.HasExtension(s.Desc.Options, persist.E_ServiceType) {
		ext, err := proto.GetExtension(s.Desc.Options, persist.E_ServiceType)
		if err == nil {
			return ext.(*persist.PersistenceOptions)
		}
	}
	return nil
}

func (s *Service) IsSQL() bool {
	if p := s.GetServiceType(); p != nil {
		if *p == persist.PersistenceOptions_SQL {
			return true
		}
	}
	return false
}

func (s *Service) IsSpanner() bool {
	if p := s.GetServiceType(); p != nil {
		if *p == persist.PersistenceOptions_SPANNER {
			return true
		}
	}
	return false
}

func (s *Service) PrintBuilder(cacheForTypeMappingNames map[string]bool) string {
	p := PersistStringer{}
	return p.PersistImplBuilder(s, cacheForTypeMappingNames)
}

type Services []*Service

// we are a persist service if we have persist options. meaning we are either spanner
// or sql
func (s Services) HasPersistService() bool {
	for _, serv := range s {
		if serv.IsSQL() || serv.IsSpanner() {
			return true
		}
	}
	return false
}

func (s *Services) AddService(pkg string, desc *descriptor.ServiceDescriptorProto, allStructs *StructList, file *FileStruct) *Service {
	ret := &Service{
		Package:    pkg,
		Desc:       desc,
		Methods:    &Methods{},
		AllStructs: allStructs,
		File:       file,
	}
	ret.ProcessMethods()
	logrus.Debugf("created a service: %s", ret)
	*s = append(*s, ret)
	return ret
}

func (s *Services) Process() error {
	for _, srv := range *s {
		err := srv.Process()
		if err != nil {
			return fmt.Errorf("%s\n  service: %s", err, srv.GetName())
		}
	}
	return nil
}

func (s *Services) PreGenerate() error {
	for _, srv := range *s {
		err := srv.Methods.PreGenerate()
		if err != nil {
			return fmt.Errorf("%s\n  service: %s", err, srv.GetName())
		}
	}
	return nil
}
