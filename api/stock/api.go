package stock

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	stock.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	stock.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
