// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: apim.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/apim/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.dop.apim",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType             = reflect.TypeOf((*Client)(nil)).Elem()
	exportRecordsClientType = reflect.TypeOf((*pb.ExportRecordsClient)(nil)).Elem()
	exportRecordsServerType = reflect.TypeOf((*pb.ExportRecordsServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.dop.apim-client":
		return p.client
	case "erda.dop.apim.ExportRecords":
		return &exportRecordsWrapper{client: p.client.ExportRecords(), opts: opts}
	case "erda.dop.apim.ExportRecords.client":
		return p.client.ExportRecords()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case exportRecordsClientType:
		return p.client.ExportRecords()
	case exportRecordsServerType:
		return &exportRecordsWrapper{client: p.client.ExportRecords(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.apim-client", &servicehub.Spec{
		Services: []string{
			"erda.dop.apim.ExportRecords",
			"erda.dop.apim.ExportRecords.client",
			"erda.dop.apim-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			exportRecordsClientType,
			// server types
			exportRecordsServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}