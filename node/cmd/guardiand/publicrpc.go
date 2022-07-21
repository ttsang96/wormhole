package guardiand

import (
	"fmt"
	"net"

	publicrpcv1 "node/proto/publicrpc/v1"

	"node/pkg/common"
	"node/pkg/db"
	"node/pkg/publicrpc"
	"node/pkg/supervisor"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func publicrpcServiceRunnable(logger *zap.Logger, listenAddr string, db *db.Database, gst *common.GuardianSetState) (supervisor.Runnable, *grpc.Server, error) {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %w", err)
	}

	logger.Info("publicrpc server listening", zap.String("addr", l.Addr().String()))

	rpcServer := publicrpc.NewPublicrpcServer(logger, db, gst)
	grpcServer := common.NewInstrumentedGRPCServer(logger)
	publicrpcv1.RegisterPublicRPCServiceServer(grpcServer, rpcServer)

	return supervisor.GRPCServer(grpcServer, l, false), grpcServer, nil
}
