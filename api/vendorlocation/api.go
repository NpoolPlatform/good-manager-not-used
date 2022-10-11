package vendorlocation

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	vendorlocation.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	vendorlocation.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
