// Copyright 2018 The issues2markdown Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package issues2markdown

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

const (
	// DefaultQuery is the default query to be used if none is provided on the
	// CLI arguments.
	DefaultQuery = `type:issue is:open author:{{ .Organization }} archived:false`
)

// QueryOptions are the available options to modify the query of issues
type QueryOptions struct {
	Organization string
}

// NewQueryOptions creates a new QueryOptions instance with sensible defaults
func NewQueryOptions() *QueryOptions {
	options := &QueryOptions{}
	return options
}

// BuildQuey builds the query string to query issues
//
// It modifies the default query according the proviced query options
func (qo *QueryOptions) BuildQuey(q string) string {
	query := strings.Builder{}

	// If query is none we use the default one
	if q == "" {
		var compiled bytes.Buffer
		t := template.Must(template.New("issueslist").Parse(DefaultQuery))
		_ = t.Execute(&compiled, qo)
		return compiled.String()
	}

	// whe only want issues
	_, _ = query.WriteString("type:issue")
	// append queries provided by CLI arguments
	_, _ = query.WriteString(fmt.Sprintf(" %s", q))
	return query.String()
}
