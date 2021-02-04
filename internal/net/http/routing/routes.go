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

// Package routing provides implementation of Go API for routing http Handler wrapped by rest.Func
package routing

import "github.com/vdaas/vald/internal/net/http/rest"

// Route struct.
type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc rest.Func
}