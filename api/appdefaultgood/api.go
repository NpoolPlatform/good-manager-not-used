package appdefaultgood

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appdefaultgood.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	appdefaultgood.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
