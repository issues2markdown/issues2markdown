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

const (
	// DefaultIssueTemplate is the default template to render a list of issues in Markdown
	DefaultIssueTemplate = `{{ range . }}- [{{ if eq .State "closed" }}x{{ else }} {{ end }}] {{ .GetOrganization }}/{{ .GetRepository }} : [#{{.Number}} {{ .Title }}]({{ .HTMLURL }})
{{ end }}`
)

// RenderOptions are the available options to modify the rendering of issues
type RenderOptions struct {
	TemplateSource string
}

// NewRenderOptions creates a new RenderOptions instance with sensible defaults
func NewRenderOptions() *RenderOptions {
	options := &RenderOptions{
		TemplateSource: DefaultIssueTemplate,
	}
	return options
}
