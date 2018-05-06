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

package issues2markdown_test

import (
	"testing"

	"github.com/issues2markdown/issues2markdown"
)

func TestIntanceQueryOptions(t *testing.T) {
	options := issues2markdown.NewQueryOptions("username")

	expectedOrganization := "username"
	if options.Organization != expectedOrganization {
		t.Fatalf("Default Organization filter expected to be %q but got %q", expectedOrganization, options.Organization)
	}

	expectedRepository := ""
	if options.Repository != expectedRepository {
		t.Fatalf("Default Repository filter expected to be %q but got %q", expectedRepository, options.Repository)
	}

	expectedState := "all"
	if options.State != expectedState {
		t.Fatalf("Default State filter expected to be %q but got %q", expectedState, options.State)
	}
}

func TestBuildQueryQueryOptions(t *testing.T) {
	options := issues2markdown.NewQueryOptions("username")

	expectedQuery := "type:issue org:username state:open state:closed"
	query := options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("Default QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.Organization = "organization"
	expectedQuery = "type:issue org:organization state:open state:closed"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.Repository = "repository"
	expectedQuery = "type:issue repo:organization/repository state:open state:closed"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.State = "open"
	expectedQuery = "type:issue repo:organization/repository state:open"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.State = "closed"
	expectedQuery = "type:issue repo:organization/repository state:closed"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}
}
