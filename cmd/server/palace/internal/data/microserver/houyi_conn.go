package microserver

import (
	"context"

	"github.com/aide-cloud/moon/api/merr"
	metadataapi "github.com/aide-cloud/moon/api/metadata"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/palaceconf"
	"github.com/aide-cloud/moon/pkg/types"
	"github.com/aide-cloud/moon/pkg/vobj"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"
)

// NewHouYiConn 创建一个HouYi rpc连接
func NewHouYiConn(c *palaceconf.Bootstrap) (*HouYiConn, func(), error) {
	microServer := c.GetMicroServer()
	houYiServer := microServer.GetHouYiServer()
	houYiConn := &HouYiConn{}
	switch houYiServer.GetNetwork() {
	case "http", "HTTP":
		httpConn, err := newHttpConn(houYiServer, microServer.GetDiscovery())
		if !types.IsNil(err) {
			log.Errorw("连接HouYi http失败：", err)
			return nil, nil, err
		}
		houYiConn.httpClient = httpConn
		houYiConn.network = vobj.NetworkHttp
	case "https", "HTTPS":
		httpConn, err := newHttpConn(houYiServer, microServer.GetDiscovery())
		if !types.IsNil(err) {
			log.Errorw("连接HouYi http失败：", err)
			return nil, nil, err
		}
		houYiConn.httpClient = httpConn
		houYiConn.network = vobj.NetworkHttps
	case "rpc", "RPC", "grpc", "GRPC":
		grpcConn, err := newRpcConn(houYiServer, microServer.GetDiscovery())
		if !types.IsNil(err) {
			log.Errorw("连接HouYi rpc失败：", err)
			return nil, nil, err
		}
		houYiConn.rpcClient = grpcConn
		houYiConn.network = vobj.NetworkRpc
	default:
		return nil, nil, merr.ErrorNotification("HouYi Server暂不支持该网络类型：[%s]", houYiServer.GetNetwork())
	}
	// 退出时清理资源
	cleanup := func() {
		if houYiConn.rpcClient != nil {
			if err := houYiConn.rpcClient.Close(); !types.IsNil(err) {
				log.Errorw("关闭 houYi rpc 连接失败：", err)
			}
		}
		if houYiConn.httpClient != nil {
			if err := houYiConn.httpClient.Close(); !types.IsNil(err) {
				log.Errorw("关闭 houYi http 连接失败：", err)
			}
		}
		log.Info("关闭 rabbit rpc连接已完成")
	}

	return houYiConn, cleanup, nil
}

type HouYiConn struct {
	// rpc连接
	rpcClient *grpc.ClientConn
	// 网络请求类型
	network vobj.Network
	// http连接
	httpClient *http.Client
}

func (l *HouYiConn) Sync(ctx context.Context, in *metadataapi.SyncRequest, opts ...Option) (*metadataapi.SyncReply, error) {
	switch l.network {
	case vobj.NetworkHttp, vobj.NetworkHttps:
		httpOpts := make([]http.CallOption, 0)
		for _, opt := range opts {
			httpOpts = append(httpOpts, opt.HttpOpts...)
		}
		return metadataapi.NewMetricHTTPClient(l.httpClient).Sync(ctx, in, httpOpts...)
	default:
		rpcOpts := make([]grpc.CallOption, 0)
		for _, opt := range opts {
			rpcOpts = append(rpcOpts, opt.RpcOpts...)
		}
		return metadataapi.NewMetricClient(l.rpcClient).Sync(ctx, in, rpcOpts...)
	}
}