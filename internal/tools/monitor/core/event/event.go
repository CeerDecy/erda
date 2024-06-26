// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package event

import "time"

// Event .
type Event struct {
	EventID   string            `json:"event_id" ch:"event_id"`
	Name      string            `json:"name" ch:"name"`
	Kind      string            `json:"kind" ch:"kind"`
	Content   string            `json:"content" ch:"content"`
	Timestamp int64             `json:"timestamp"`
	Time      time.Time         `json:"-" ch:"timestamp"`
	Tags      map[string]string `json:"tags" ch:"tags"`
	Relations map[string]string `json:"relations" ch:"relations"`
}
