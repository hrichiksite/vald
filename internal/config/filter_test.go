//
// Copyright (C) 2019-2021 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package config providers configuration type and load configuration logic
package config

import (
	"reflect"
	"testing"

	"github.com/vdaas/vald/internal/errors"
	"go.uber.org/goleak"
)

func TestEgressFilter_Bind(t *testing.T) {
	t.Parallel()
	type fields struct {
		Client          *GRPCClient
		DistanceFilters []string
		ObjectFilters   []string
	}
	type want struct {
		want *EgressFilter
	}
	type test struct {
		name       string
		fields     fields
		want       want
		checkFunc  func(want, *EgressFilter) error
		beforeFunc func()
		afterFunc  func()
	}
	defaultCheckFunc := func(w want, got *EgressFilter) error {
		if !reflect.DeepEqual(got, w.want) {
			return errors.Errorf("got: \"%#v\",\n\t\t\t\twant: \"%#v\"", got, w.want)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       fields: fields {
		           Client: GRPCClient{},
		           DistanceFilters: nil,
		           ObjectFilters: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           fields: fields {
		           Client: GRPCClient{},
		           DistanceFilters: nil,
		           ObjectFilters: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			defer goleak.VerifyNone(tt, goleak.IgnoreCurrent())
			if test.beforeFunc != nil {
				test.beforeFunc()
			}
			if test.afterFunc != nil {
				defer test.afterFunc()
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			e := &EgressFilter{
				Client:          test.fields.Client,
				DistanceFilters: test.fields.DistanceFilters,
				ObjectFilters:   test.fields.ObjectFilters,
			}

			got := e.Bind()
			if err := test.checkFunc(test.want, got); err != nil {
				tt.Errorf("error = %v", err)
			}
		})
	}
}

func TestIngressFilter_Bind(t *testing.T) {
	t.Parallel()
	type fields struct {
		Client        *GRPCClient
		Vectorizer    string
		SearchFilters []string
		InsertFilters []string
		UpdateFilters []string
		UpsertFilters []string
	}
	type want struct {
		want *IngressFilter
	}
	type test struct {
		name       string
		fields     fields
		want       want
		checkFunc  func(want, *IngressFilter) error
		beforeFunc func()
		afterFunc  func()
	}
	defaultCheckFunc := func(w want, got *IngressFilter) error {
		if !reflect.DeepEqual(got, w.want) {
			return errors.Errorf("got: \"%#v\",\n\t\t\t\twant: \"%#v\"", got, w.want)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       fields: fields {
		           Client: GRPCClient{},
		           Vectorizer: "",
		           SearchFilters: nil,
		           InsertFilters: nil,
		           UpdateFilters: nil,
		           UpsertFilters: nil,
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           fields: fields {
		           Client: GRPCClient{},
		           Vectorizer: "",
		           SearchFilters: nil,
		           InsertFilters: nil,
		           UpdateFilters: nil,
		           UpsertFilters: nil,
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			defer goleak.VerifyNone(tt, goleak.IgnoreCurrent())
			if test.beforeFunc != nil {
				test.beforeFunc()
			}
			if test.afterFunc != nil {
				defer test.afterFunc()
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			i := &IngressFilter{
				Client:        test.fields.Client,
				Vectorizer:    test.fields.Vectorizer,
				SearchFilters: test.fields.SearchFilters,
				InsertFilters: test.fields.InsertFilters,
				UpdateFilters: test.fields.UpdateFilters,
				UpsertFilters: test.fields.UpsertFilters,
			}

			got := i.Bind()
			if err := test.checkFunc(test.want, got); err != nil {
				tt.Errorf("error = %v", err)
			}
		})
	}
}