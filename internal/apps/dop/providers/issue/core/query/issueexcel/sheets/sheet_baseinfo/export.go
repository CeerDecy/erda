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

package sheet_baseinfo

import (
	"encoding/json"
	"fmt"

	"github.com/erda-project/erda/internal/apps/dop/conf"
	"github.com/erda-project/erda/internal/apps/dop/providers/issue/core/query/issueexcel/sheets"
	"github.com/erda-project/erda/internal/apps/dop/providers/issue/core/query/issueexcel/vars"
	"github.com/erda-project/erda/pkg/excel"
)

func (h *Handler) ExportSheet(data *vars.DataForFulfill) (*sheets.RowsForExport, error) {
	// only one row, k=meta, v=JSON(dataForFulfillImportOnlyBaseInfo)
	meta := vars.DataForFulfillImportOnlyBaseInfo{
		OriginalErdaPlatform:  conf.DiceClusterName(),
		OriginalErdaProjectID: data.ProjectID,
		AllProjectIssues:      data.ExportOnly.IsFullExport,
	}
	b, err := json.Marshal(&meta)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal meta, err: %v", err)
	}
	var row excel.Row
	row = append(row, excel.NewCell(cellMeta))
	row = append(row, excel.NewCell(string(b)))
	return sheets.NewRowsForExport(excel.Rows{row}), nil
}
