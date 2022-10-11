package appgood

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appgood.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appgood.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
