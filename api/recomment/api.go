package recommend

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	recommend.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	recommend.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
