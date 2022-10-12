//nolint:nolintlint,dupl
package deviceinfo

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"

	constant "github.com/NpoolPlatform/good-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateDeviceInfo(ctx context.Context, in *npool.DeviceInfoReq) (*npool.DeviceInfo, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateDeviceInfo(ctx, &npool.CreateDeviceInfoRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}

func CreateDeviceInfos(ctx context.Context, in []*npool.DeviceInfoReq) ([]*npool.DeviceInfo, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateDeviceInfos(ctx, &npool.CreateDeviceInfosRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.DeviceInfo), nil
}

func UpdateDeviceInfo(ctx context.Context, in *npool.DeviceInfoReq) (*npool.DeviceInfo, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateDeviceInfo(ctx, &npool.UpdateDeviceInfoRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}

func GetDeviceInfo(ctx context.Context, id string) (*npool.DeviceInfo, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetDeviceInfo(ctx, &npool.GetDeviceInfoRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}

func GetDeviceInfoOnly(ctx context.Context, conds *npool.Conds) (*npool.DeviceInfo, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetDeviceInfoOnly(ctx, &npool.GetDeviceInfoOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}

func GetDeviceInfos(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.DeviceInfo, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetDeviceInfos(ctx, &npool.GetDeviceInfosRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.DeviceInfo), total, nil
}

func ExistDeviceInfo(ctx context.Context, id string) (bool, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistDeviceInfo(ctx, &npool.ExistDeviceInfoRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func ExistDeviceInfoConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistDeviceInfoConds(ctx, &npool.ExistDeviceInfoCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func DeleteDeviceInfo(ctx context.Context, id string) (*npool.DeviceInfo, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteDeviceInfo(ctx, &npool.DeleteDeviceInfoRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}
