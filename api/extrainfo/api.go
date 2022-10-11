package extrainfo

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/extrainfo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	extrainfo.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	extrainfo.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
