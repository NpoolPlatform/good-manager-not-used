package deviceinfo

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	deviceinfo.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	deviceinfo.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
