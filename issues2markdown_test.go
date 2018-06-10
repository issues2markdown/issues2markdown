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
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/github"
	"github.com/issues2markdown/issues2markdown"
	"golang.org/x/oauth2"
)

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

// providerSetup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server.
//
// Tests must register handlers on mux which provide mock responses for the
// API method being tested.
func providerSetup(t *testing.T) (client *github.Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	ctx := context.Background()

	// client is the GitHub client being tested and is
	// configured to use test server.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: "github_token",
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)

	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url
	client.UploadURL = url

	return client, mux, server.URL, server.Close
}

func TestInstanceIssuesToMarkdown(t *testing.T) {
	issuesProvider, mux, _, teardown := providerSetup(t)
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprint(w, `{"login": "username"}`)
	})
	defer teardown()

	_, err := issues2markdown.NewIssuesToMarkdown(issuesProvider)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInstanceIssuesToMarkdownUnauthorized(t *testing.T) {
	issuesProvider, mux, _, teardown := providerSetup(t)
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		http.Error(w, "Unauthorized", 401)
	})
	defer teardown()

	_, err := issues2markdown.NewIssuesToMarkdown(issuesProvider)
	if err == nil {
		t.Fatal(err)
	}
}

func TestQuery(t *testing.T) {
	issuesProvider, mux, _, teardown := providerSetup(t)
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprint(w, `{"login": "username"}`)
	})
	mux.HandleFunc("/search/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprint(w, `{}`)
	})
	defer teardown()

	i2md, err := issues2markdown.NewIssuesToMarkdown(issuesProvider)
	if err != nil {
		t.Fatal(err)
	}

	q := ""
	options := issues2markdown.NewQueryOptions()
	_, err = i2md.Query(options, q)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRender(t *testing.T) {
	issuesProvider, mux, _, teardown := providerSetup(t)
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprint(w, `{"login": "username"}`)
	})
	issues := []issues2markdown.Issue{
		{
			Number:  1,
			Title:   "Issue title 1",
			State:   "open",
			URL:     "https://api.github.com/repos/username/repo/issues/1",
			HTMLURL: "https://github.com/username/repo/issues/1",
		},
		{
			Number:  2,
			Title:   "Issue title 2",
			State:   "closed",
			URL:     "https://api.github.com/repos/username/repo/issues/2",
			HTMLURL: "https://github.com/username/repo/issues/2",
		},
	}
	defer teardown()

	i2md, err := issues2markdown.NewIssuesToMarkdown(issuesProvider)
	if err != nil {
		t.Fatal(err)
	}

	options := issues2markdown.NewRenderOptions()
	markdown, err := i2md.Render(issues, options)
	if err != nil {
		t.Fatal(err)
	}

	expectedMarkdown := `- [ ] username/repo : [#1 Issue title 1](https://github.com/username/repo/issues/1)
- [x] username/repo : [#2 Issue title 2](https://github.com/username/repo/issues/2)`

	if markdown != expectedMarkdown {
		t.Fatalf("Expected %q but got %q", expectedMarkdown, markdown)
	}
}
