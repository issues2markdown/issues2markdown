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
	"context"
	"html/template"
	"strings"

	"github.com/google/go-github/github"
)

// IssuesToMarkdown is the main type to interact, query and render issues to
// Markdown
type IssuesToMarkdown struct {
	GithubToken string
	client      *github.Client
	User        *github.User
}

// NewIssuesToMarkdown creates an IssuesToMarkdown instance
func NewIssuesToMarkdown(provider *github.Client) (*IssuesToMarkdown, error) {
	i2md := &IssuesToMarkdown{
		client: provider,
	}
	user, err := i2md.Authorize()
	if err != nil {
		return nil, err
	}
	i2md.User = user
	return i2md, nil
}

// Authorize gets authentication information
func (im *IssuesToMarkdown) Authorize() (*github.User, error) {
	ctx := context.Background()
	// get user information
	user, _, err := im.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	return user, err
}

// Query queries the provider and returns the list of Issues that match
// the query
func (im *IssuesToMarkdown) Query(options *QueryOptions, q string) ([]Issue, error) {
	ctx := context.Background()

	// query issues
	var result []Issue
	query := options.BuildQuey(q)

	githubOptions := &github.SearchOptions{}
	for {
		listResult, response, err := im.client.Search.Issues(ctx, query, githubOptions)
		if err != nil {
			return nil, err
		}

		// process page results
		for _, v := range listResult.Issues {
			item := Issue{
				Number:  *v.Number,
				Title:   *v.Title,
				State:   *v.State,
				URL:     *v.URL,
				HTMLURL: *v.HTMLURL,
			}
			result = append(result, item)
		}

		// process pagination
		if response.NextPage == 0 {
			break
		}
		githubOptions.Page = response.NextPage
	}

	return result, nil
}

// Render renders a list of Issues to Markdown
func (im *IssuesToMarkdown) Render(issues []Issue, options *RenderOptions) (string, error) {
	var compiled bytes.Buffer
	t := template.Must(template.New("issueslist").Parse(options.TemplateSource))
	_ = t.Execute(&compiled, issues)
	result := compiled.String()
	result = strings.TrimRight(result, "\n") // trim the last linebreak
	return result, nil
}
