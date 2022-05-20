//
// Copyright (C) 2019-2022 vdaas.org vald team <vald@vdaas.org>
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
package operation

import (
	"context"
	"reflect"
	"testing"

	"github.com/vdaas/vald/hack/benchmark/internal/assets"
	"github.com/vdaas/vald/internal/client/v1/client"
	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/internal/test/goleak"
)

func Test_operation_Insert(t *testing.T) {
	type args struct {
		ctx context.Context
		b   *testing.B
		ds  assets.Dataset
	}
	type fields struct {
		client   client.Client
		indexerC client.Indexer
	}
	type want struct {
		wantInsertedNum int
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, gotInsertedNum int) error {
		if !reflect.DeepEqual(gotInsertedNum, w.wantInsertedNum) {
			return errors.Errorf("got: \"%#v\",\n\t\t\t\twant: \"%#v\"", gotInsertedNum, w.wantInsertedNum)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       args: args {
		           ctx: nil,
		           b: testing.B{},
		           ds: nil,
		       },
		       fields: fields {
		           client: nil,
		           indexerC: nil,
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
		           args: args {
		           ctx: nil,
		           b: testing.B{},
		           ds: nil,
		           },
		           fields: fields {
		           client: nil,
		           indexerC: nil,
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
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			o := &operation{
				client:   test.fields.client,
				indexerC: test.fields.indexerC,
			}

			gotInsertedNum := o.Insert(test.args.ctx, test.args.b, test.args.ds)
			if err := test.checkFunc(test.want, gotInsertedNum); err != nil {
				tt.Errorf("error = %v", err)
			}
		})
	}
}

func Test_operation_StreamInsert(t *testing.T) {
	type args struct {
		ctx context.Context
		b   *testing.B
		ds  assets.Dataset
	}
	type fields struct {
		client   client.Client
		indexerC client.Indexer
	}
	type want struct {
		want int
	}
	type test struct {
		name       string
		args       args
		fields     fields
		want       want
		checkFunc  func(want, int) error
		beforeFunc func(args)
		afterFunc  func(args)
	}
	defaultCheckFunc := func(w want, got int) error {
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
		       args: args {
		           ctx: nil,
		           b: testing.B{},
		           ds: nil,
		       },
		       fields: fields {
		           client: nil,
		           indexerC: nil,
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
		           args: args {
		           ctx: nil,
		           b: testing.B{},
		           ds: nil,
		           },
		           fields: fields {
		           client: nil,
		           indexerC: nil,
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
				test.beforeFunc(test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(test.args)
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			o := &operation{
				client:   test.fields.client,
				indexerC: test.fields.indexerC,
			}

			got := o.StreamInsert(test.args.ctx, test.args.b, test.args.ds)
			if err := test.checkFunc(test.want, got); err != nil {
				tt.Errorf("error = %v", err)
			}
		})
	}
}