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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: apis/proto/v1/vald/filter.proto

package vald

import (
	reflect "reflect"

	payload "github.com/vdaas/vald/apis/grpc/v1/payload"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_apis_proto_v1_vald_filter_proto protoreflect.FileDescriptor

var file_apis_proto_v1_vald_filter_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x61, 0x6c, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x76, 0x61, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xac,
	0x0a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0c, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x0e, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x7c, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x5f, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x68, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x69, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x12,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x7c, 0x0a,
	0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22,
	0x17, 0x2f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x7c, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x75, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5f,
	0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x7c, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x22, 0x17, 0x2f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x53, 0x0a,
	0x1a, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x64, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x61, 0x6c, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x61, 0x6c, 0x64, 0x42, 0x0a, 0x56, 0x61, 0x6c,
	0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x64, 0x61, 0x61, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61,
	0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apis_proto_v1_vald_filter_proto_goTypes = []interface{}{
	(*payload.Search_ObjectRequest)(nil),      // 0: payload.v1.Search.ObjectRequest
	(*payload.Search_MultiObjectRequest)(nil), // 1: payload.v1.Search.MultiObjectRequest
	(*payload.Insert_ObjectRequest)(nil),      // 2: payload.v1.Insert.ObjectRequest
	(*payload.Insert_MultiObjectRequest)(nil), // 3: payload.v1.Insert.MultiObjectRequest
	(*payload.Update_ObjectRequest)(nil),      // 4: payload.v1.Update.ObjectRequest
	(*payload.Update_MultiObjectRequest)(nil), // 5: payload.v1.Update.MultiObjectRequest
	(*payload.Upsert_ObjectRequest)(nil),      // 6: payload.v1.Upsert.ObjectRequest
	(*payload.Upsert_MultiObjectRequest)(nil), // 7: payload.v1.Upsert.MultiObjectRequest
	(*payload.Search_Response)(nil),           // 8: payload.v1.Search.Response
	(*payload.Search_Responses)(nil),          // 9: payload.v1.Search.Responses
	(*payload.Search_StreamResponse)(nil),     // 10: payload.v1.Search.StreamResponse
	(*payload.Object_Location)(nil),           // 11: payload.v1.Object.Location
	(*payload.Object_StreamLocation)(nil),     // 12: payload.v1.Object.StreamLocation
	(*payload.Object_Locations)(nil),          // 13: payload.v1.Object.Locations
}
var file_apis_proto_v1_vald_filter_proto_depIdxs = []int32{
	0,  // 0: vald.v1.Filter.SearchObject:input_type -> payload.v1.Search.ObjectRequest
	1,  // 1: vald.v1.Filter.MultiSearchObject:input_type -> payload.v1.Search.MultiObjectRequest
	0,  // 2: vald.v1.Filter.StreamSearchObject:input_type -> payload.v1.Search.ObjectRequest
	2,  // 3: vald.v1.Filter.InsertObject:input_type -> payload.v1.Insert.ObjectRequest
	2,  // 4: vald.v1.Filter.StreamInsertObject:input_type -> payload.v1.Insert.ObjectRequest
	3,  // 5: vald.v1.Filter.MultiInsertObject:input_type -> payload.v1.Insert.MultiObjectRequest
	4,  // 6: vald.v1.Filter.UpdateObject:input_type -> payload.v1.Update.ObjectRequest
	4,  // 7: vald.v1.Filter.StreamUpdateObject:input_type -> payload.v1.Update.ObjectRequest
	5,  // 8: vald.v1.Filter.MultiUpdateObject:input_type -> payload.v1.Update.MultiObjectRequest
	6,  // 9: vald.v1.Filter.UpsertObject:input_type -> payload.v1.Upsert.ObjectRequest
	6,  // 10: vald.v1.Filter.StreamUpsertObject:input_type -> payload.v1.Upsert.ObjectRequest
	7,  // 11: vald.v1.Filter.MultiUpsertObject:input_type -> payload.v1.Upsert.MultiObjectRequest
	8,  // 12: vald.v1.Filter.SearchObject:output_type -> payload.v1.Search.Response
	9,  // 13: vald.v1.Filter.MultiSearchObject:output_type -> payload.v1.Search.Responses
	10, // 14: vald.v1.Filter.StreamSearchObject:output_type -> payload.v1.Search.StreamResponse
	11, // 15: vald.v1.Filter.InsertObject:output_type -> payload.v1.Object.Location
	12, // 16: vald.v1.Filter.StreamInsertObject:output_type -> payload.v1.Object.StreamLocation
	13, // 17: vald.v1.Filter.MultiInsertObject:output_type -> payload.v1.Object.Locations
	11, // 18: vald.v1.Filter.UpdateObject:output_type -> payload.v1.Object.Location
	12, // 19: vald.v1.Filter.StreamUpdateObject:output_type -> payload.v1.Object.StreamLocation
	13, // 20: vald.v1.Filter.MultiUpdateObject:output_type -> payload.v1.Object.Locations
	11, // 21: vald.v1.Filter.UpsertObject:output_type -> payload.v1.Object.Location
	12, // 22: vald.v1.Filter.StreamUpsertObject:output_type -> payload.v1.Object.StreamLocation
	13, // 23: vald.v1.Filter.MultiUpsertObject:output_type -> payload.v1.Object.Locations
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_apis_proto_v1_vald_filter_proto_init() }
func file_apis_proto_v1_vald_filter_proto_init() {
	if File_apis_proto_v1_vald_filter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_proto_v1_vald_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_proto_v1_vald_filter_proto_goTypes,
		DependencyIndexes: file_apis_proto_v1_vald_filter_proto_depIdxs,
	}.Build()
	File_apis_proto_v1_vald_filter_proto = out.File
	file_apis_proto_v1_vald_filter_proto_rawDesc = nil
	file_apis_proto_v1_vald_filter_proto_goTypes = nil
	file_apis_proto_v1_vald_filter_proto_depIdxs = nil
}