// Copyright 2018 Lyceum Developers
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

package models

// Organization represents a single tenant or project to segment libraries
type Organization struct {
	DateCreated  string   `json:"date_created"`
	DateModified string   `json:"date_modified"`
	Description  string   `json:"description"`
	Name         string   `json:"name"`
	ID           string   `json:"id"`
	Libraries    []string `json:"libraries"`
	Users        []string `json:"users"`
}
