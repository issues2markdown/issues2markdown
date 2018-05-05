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
	"net/url"
	"strings"
)

// Issue represents an Issue from the provider
type Issue struct {
	Title   string
	State   string
	URL     string
	HTMLURL string
}

// NewIssue creates an Issue instance with sensible defaults
func NewIssue() *Issue {
	issue := &Issue{}
	return issue
}

// GetOrganization return the organization name for this Issue
func (i *Issue) GetOrganization() (string, error) {
	parsedU, _ := url.Parse(i.URL)
	parsedPartsPathU := strings.Split(parsedU.Path, "/")
	organization := parsedPartsPathU[2]
	return organization, nil
}

// GetRepository return the repository name for this Issue
func (i *Issue) GetRepository() (string, error) {
	parsedU, _ := url.Parse(i.URL)
	parsedPartsPathU := strings.Split(parsedU.Path, "/")
	repository := parsedPartsPathU[3]
	return repository, nil
}
