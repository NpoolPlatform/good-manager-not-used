package subgood

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/subgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	subgood.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	subgood.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
