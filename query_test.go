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
	options := issues2markdown.NewQueryOptions()
	options.Organization = "username"

	expectedOrganization := "username"
	if options.Organization != expectedOrganization {
		t.Fatalf("Default Organization filter expected to be %q but got %q", expectedOrganization, options.Organization)
	}
}

func TestBuildQueryDefaultQueryOptions(t *testing.T) {
	options := issues2markdown.NewQueryOptions()
	options.Organization = "username"

	expectedQuery := "type:issue is:open author:username archived:false"
	query := options.BuildQuey("")
	if query != expectedQuery {
		t.Fatalf("Default QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}
}

func TestBuildQueryQueryOptions(t *testing.T) {
	options := issues2markdown.NewQueryOptions()
	options.Organization = "username"

	expectedQuery := "type:issue repo:organization/repository"
	query := options.BuildQuey("repo:organization/repository")
	if query != expectedQuery {
		t.Fatalf("Default QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}
}
