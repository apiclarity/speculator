// Copyright © 2021 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import "github.com/getkin/kin-openapi/openapi3"

type LearningSpec struct {
	// map parameterized path into path item
	PathItems       map[string]*openapi3.PathItem
	SecuritySchemes openapi3.SecuritySchemes
}

func (l *LearningSpec) AddPathItem(path string, pathItem *openapi3.PathItem) {
	l.PathItems[path] = pathItem
}

func (l *LearningSpec) GetPathItem(path string) *openapi3.PathItem {
	pi, ok := l.PathItems[path]
	if !ok {
		return nil
	}

	return pi
}
