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

package task

import (
	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/pkg/transport"
	"github.com/erda-project/erda-infra/providers/mysqlxorm"
	"github.com/erda-project/erda-proto-go/core/pipeline/task/pb"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/internal/tools/pipeline/dbclient"
	"github.com/erda-project/erda/internal/tools/pipeline/providers/edgepipeline_register"
	"github.com/erda-project/erda/internal/tools/pipeline/providers/permission"
	"github.com/erda-project/erda/internal/tools/pipeline/providers/pipeline"
	"github.com/erda-project/erda/pkg/common/apis"
)

type config struct {
}

// +provider
type provider struct {
	Cfg          *config
	Log          logs.Logger
	Register     transport.Register
	MySQL        mysqlxorm.Interface
	PipelineSvc  pipeline.Interface
	Permission   permission.Interface
	EdgeRegister edgepipeline_register.Interface

	taskService *taskService
}

func (p *provider) Init(ctx servicehub.Context) error {
	bdl := bundle.New(bundle.WithErdaServer())
	p.taskService = &taskService{
		p:            p,
		dbClient:     &dbclient.Client{Engine: p.MySQL.DB()},
		bdl:          bdl,
		pipelineSvc:  p.PipelineSvc,
		permission:   p.Permission,
		edgeRegister: p.EdgeRegister,
	}
	if p.Register != nil {
		pb.RegisterTaskServiceImp(p.Register, p.taskService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.pipeline.task.TaskService" || ctx.Type() == pb.TaskServiceServerType() || ctx.Type() == pb.TaskServiceHandlerType():
		return p.taskService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.pipeline.task", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Description:          "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
