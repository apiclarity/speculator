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

package utils

import "testing"

func TestIsApplicationJsonMediaType(t *testing.T) {
	type args struct {
		mediaType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "application/json",
			args: args{
				mediaType: "application/json",
			},
			want: true,
		},
		{
			name: "application/hal+json",
			args: args{
				mediaType: "application/hal+json",
			},
			want: true,
		},
		{
			name: "not application json mime",
			args: args{
				mediaType: "test/html",
			},
			want: false,
		},
		{
			name: "empty mediaType",
			args: args{
				mediaType: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsApplicationJSONMediaType(tt.args.mediaType); got != tt.want {
				t.Errorf("IsApplicationJSONMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}
