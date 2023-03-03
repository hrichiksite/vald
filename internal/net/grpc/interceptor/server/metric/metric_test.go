// Copyright (C) 2019-2023 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package metric

import (
	"context"
	"reflect"
	"testing"

	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/internal/net/grpc/codes"
	"github.com/vdaas/vald/internal/net/grpc/status"
	"github.com/vdaas/vald/internal/observability/attribute"
	"github.com/vdaas/vald/internal/test/goleak"
)

func Test_attributesFromError(t *testing.T) {
	type T = []attribute.KeyValue
	type args struct {
		method string
		err    error
	}
	type want struct {
		obj T
	}
	type test struct {
		name       string
		args       args
		want       want
		checkFunc  func(w want, obj T) error
		beforeFunc func(*testing.T, args)
		afterFunc  func(*testing.T, args)
	}

	defaultCheckFunc := func(w want, obj T) error {
		if !reflect.DeepEqual(obj, w.obj) {
			return errors.Errorf("got: \"%#v\",\n\t\t\t\twant: \"%#v\"", obj, w.obj)
		}
		return nil
	}

	tests := []test{
		{
			name: "return []attribute.KeyValue when err is nil",
			args: args{
				method: "InsertRPC",
				err:    nil,
			},
			want: want{
				obj: []attribute.KeyValue{
					attribute.String(gRPCMethodKeyName, "InsertRPC"),
					attribute.String(gRPCStatus, codes.OK.String()),
				},
			},
		},
		{
			name: "return []attribute.KeyValue when err is not nil",
			args: args{
				method: "InsertRPC",
				err:    context.DeadlineExceeded,
			},
			want: want{
				obj: []attribute.KeyValue{
					attribute.String(gRPCMethodKeyName, "InsertRPC"),
					attribute.String(gRPCStatus, codes.DeadlineExceeded.String()),
				},
			},
		},
		{
			name: "return []attribute.KeyValue when err type is wrapped error",
			args: args{
				method: "InsertRPC",
				err:    status.WrapWithInvalidArgument("Insert API failed", errors.ErrIncompatibleDimensionSize(100, 940)),
			},
			want: want{
				obj: []attribute.KeyValue{
					attribute.String(gRPCMethodKeyName, "InsertRPC"),
					attribute.String(gRPCStatus, codes.InvalidArgument.String()),
				},
			},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			defer goleak.VerifyNone(tt, goleak.IgnoreCurrent())
			if test.beforeFunc != nil {
				test.beforeFunc(tt, test.args)
			}
			if test.afterFunc != nil {
				defer test.afterFunc(tt, test.args)
			}

			checkFunc := test.checkFunc
			if test.checkFunc == nil {
				checkFunc = defaultCheckFunc
			}

			got := attributesFromError(test.args.method, test.args.err)
			if err := checkFunc(test.want, got); err != nil {
				tt.Errorf("error = %v", err)
			}
		})
	}
}