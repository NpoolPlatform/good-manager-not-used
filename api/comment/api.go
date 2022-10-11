package comment

import (
	"github.com/NpoolPlatform/message/npool/good/mgr/v1/comment"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	comment.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	comment.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
