package api

import (
	"context"

	"github.com/NpoolPlatform/good-manager/api/appdefaultgood"

	"github.com/NpoolPlatform/good-manager/api/appgood"
	"github.com/NpoolPlatform/good-manager/api/comment"
	"github.com/NpoolPlatform/good-manager/api/deviceinfo"
	"github.com/NpoolPlatform/good-manager/api/extrainfo"
	"github.com/NpoolPlatform/good-manager/api/good"
	"github.com/NpoolPlatform/good-manager/api/promotion"
	"github.com/NpoolPlatform/good-manager/api/recommend"
	"github.com/NpoolPlatform/good-manager/api/stock"
	"github.com/NpoolPlatform/good-manager/api/subgood"
	"github.com/NpoolPlatform/good-manager/api/vendorlocation"
	v1 "github.com/NpoolPlatform/message/npool/good/mgr/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterManagerServer(server, &Server{})
	appgood.Register(server)
	comment.Register(server)
	deviceinfo.Register(server)
	extrainfo.Register(server)
	good.Register(server)
	promotion.Register(server)
	recommend.Register(server)
	stock.Register(server)
	subgood.Register(server)
	vendorlocation.Register(server)
	appdefaultgood.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
