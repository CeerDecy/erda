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

package table

import (
	"strings"

	"github.com/erda-project/erda/internal/tools/monitor/core/settings/retention-strategy"
)

const (
	TableNameKey          = "<table_name>"
	AliasTableNameKey     = "<alias_table_name>"
	DatabaseNameKey       = "<database>"
	TtlDaysNameKey        = "<ttl_in_days>"
	TtlHotDataDaysNameKey = "<ttl_in_hot_days>"
)

var keyReplacer = strings.NewReplacer(
	"-", "_",
	".", "_",
)

// NormalizeKey .
func NormalizeKey(s string) string {
	return keyReplacer.Replace(strings.ToLower(s))
}

func FormatTTLToDays(ttl *retention.TTL) int64 {
	return ttl.GetTTLByDays()
}
