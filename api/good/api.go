package good

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	good.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	good.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
