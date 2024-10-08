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

package apistructs

// PipelineLabelType defines the type of pipeline label.
type PipelineLabelType string

var (
	PipelineLabelTypeInstance PipelineLabelType = "p_i"   // pipeline instance
	PipelineLabelTypeQueue    PipelineLabelType = "queue" // queue
)

func (t PipelineLabelType) String() string { return string(t) }
func (t PipelineLabelType) Valid() bool {
	switch t {
	case PipelineLabelTypeInstance, PipelineLabelTypeQueue:
		return true
	default:
		return false
	}
}

// TargetIDSelectByLabelRequest select target ids by labels.
type TargetIDSelectByLabelRequest struct {
	Type PipelineLabelType `json:"type"`

	PipelineSources  []PipelineSource `json:"pipelineSource"`
	PipelineYmlNames []string         `json:"pipelineYmlName"`

	// AllowNoMatchLabels, default is false.
	AllowNoMatchLabels bool `json:"allowNoMatchLabels,omitempty"`
	// MUST match
	MustMatchLabels map[string][]string `json:"mustMatchLabels"`
	// ANY match
	AnyMatchLabels map[string][]string `json:"anyMatchLabels"`

	// AllowNoPipelineSources, default is false.
	// 默认查询必须带上 pipeline source，增加区分度
	AllowNoPipelineSources bool `json:"allowNoPipelineSources"`

	// OrderByTargetIDAsc 根据 target_id 升序，默认为 false，即降序
	OrderByTargetIDAsc bool `json:"orderByTargetIDAsc"`
}
