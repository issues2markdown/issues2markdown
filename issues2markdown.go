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
	"log"
	"strings"

	"github.com/google/go-github/github"
)

// IssuesToMarkdown is the main type to interact, query and render issues to
// Markdown
type IssuesToMarkdown struct {
	client      *github.Client
	GithubToken string
	Username    string
}

// NewIssuesToMarkdown creates an IssuesToMarkdown instance and gets
// authentication information
func NewIssuesToMarkdown(provider *github.Client) (*IssuesToMarkdown, error) {
	i2md := &IssuesToMarkdown{
		client: provider,
	}

	ctx := context.Background()

	// get user information
	user, _, err := i2md.client.Users.Get(ctx, "")
	if err != nil {
		log.Printf("ERROR: %s", err)
		return nil, err
	}
	i2md.Username = user.GetLogin()
	log.Printf("Created authenticated github API client for user: %s\n", i2md.Username)

	return i2md, nil
}

// Query queries the provider and returns the list of Issues that match
// the query
func (im *IssuesToMarkdown) Query(options *QueryOptions, q string) ([]Issue, error) {
	ctx := context.Background()

	// query issues
	var result []Issue
	query := options.BuildQuey(q)
	log.Printf("Search Query: %s\n", query)

	githubOptions := &github.SearchOptions{}
	for {
		listResult, response, err := im.client.Search.Issues(ctx, query, githubOptions)
		if err != nil {
			log.Printf("ERROR: %s", err)
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
