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

package handler_audit

import (
	"context"

	"github.com/erda-project/erda-proto-go/apps/aiproxy/audit/pb"
	"github.com/erda-project/erda/internal/apps/ai-proxy/providers/dao"
)

type AuditHandler struct {
	DAO dao.DAO
}

func (a *AuditHandler) Get(ctx context.Context, req *pb.AuditGetRequest) (*pb.Audit, error) {
	return a.DAO.AuditClient().Get(ctx, req)
}

func (a *AuditHandler) Paging(ctx context.Context, req *pb.AuditPagingRequest) (*pb.AuditPagingResponse, error) {
	return a.DAO.AuditClient().Paging(ctx, req)
}