//
// Copyright (C) 2019-2023 vdaas.org vald team <vald@vdaas.org>
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

// EgressFilter represents the EgressFilter configuration.
type EgressFilter struct {
	Client          *GRPCClient             `json:"client,omitempty"           yaml:"client"`
	DistanceFilters []*DistanceFilterConfig `json:"distance_filters,omitempty" yaml:"distance_filters"`
	ObjectFilters   []*ObjectFilterConfig   `json:"object_filters,omitempty"   yaml:"object_filters"`
}

// IngressFilter represents the IngressFilter configuration.
type IngressFilter struct {
	Client        *GRPCClient `json:"client,omitempty"         yaml:"client"`
	Vectorizer    string      `json:"vectorizer,omitempty"     yaml:"vectorizer"`
	SearchFilters []string    `json:"search_filters,omitempty" yaml:"search_filters"`
	InsertFilters []string    `json:"insert_filters,omitempty" yaml:"insert_filters"`
	UpdateFilters []string    `json:"update_filters,omitempty" yaml:"update_filters"`
	UpsertFilters []string    `json:"upsert_filters,omitempty" yaml:"upsert_filters"`
}

// DistanceFilterConfig represents the DistanceFilter configuration.
type DistanceFilterConfig struct {
	Addr  string `json:"addr,omitempty"  yaml:"addr"`
	Query string `json:"query,omitempty" yaml:"query"`
}

// ObjectFilterConfig represents the ObjectFilter configuration.
type ObjectFilterConfig struct {
	Addr  string `json:"addr,omitempty"  yaml:"addr"`
	Query string `json:"query,omitempty" yaml:"query"`
}

// Bind binds the actual data from the EgressFilter receiver field.
func (e *EgressFilter) Bind() *EgressFilter {
	if e.Client != nil {
		e.Client.Bind()
	}
	for _, df := range e.DistanceFilters {
		df.Bind()
	}
	for _, of := range e.ObjectFilters {
		of.Bind()
	}
	return e
}

// Bind binds the actual data from the IngressFilter receiver field.
func (i *IngressFilter) Bind() *IngressFilter {
	if i.Client != nil {
		i.Client.Bind()
	}

	i.Vectorizer = GetActualValue(i.Vectorizer)

	if i.SearchFilters != nil {
		i.SearchFilters = GetActualValues(i.SearchFilters)
	}
	if i.InsertFilters != nil {
		i.InsertFilters = GetActualValues(i.InsertFilters)
	}
	if i.UpdateFilters != nil {
		i.UpdateFilters = GetActualValues(i.UpdateFilters)
	}
	if i.UpsertFilters != nil {
		i.UpsertFilters = GetActualValues(i.UpsertFilters)
	}
	return i
}

// Bind binds the actual data from the DistanceFilterConfig receiver field.
func (c *DistanceFilterConfig) Bind() *DistanceFilterConfig {
	c.Addr = GetActualValue(c.Addr)
	c.Query = GetActualValue(c.Query)
	return c
}

// Bind binds the actual data from the ObjectFilterConfig receiver field.
func (c *ObjectFilterConfig) Bind() *ObjectFilterConfig {
	c.Addr = GetActualValue(c.Addr)
	c.Query = GetActualValue(c.Query)
	return c
}
