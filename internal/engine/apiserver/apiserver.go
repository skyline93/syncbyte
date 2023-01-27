package apiserver

import (
	"context"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/backup"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/scheduling"
	pb "github.com/skyline93/syncbyte-go/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type apiServer struct {
	pb.UnimplementedApiServiceServer
}

func (api *apiServer) CreateBackupPolicy(ctx context.Context, req *pb.CreateBackupPolicyRequest) (res *pb.CreateBackupPolicyResponse, err error) {
	var args interface{}
	if req.ResourceType == string(backup.NAS) {
		args = backup.NasResourceArgs{Dir: req.BackupPath}
	}

	resource := backup.Resource{
		Name: req.ResourceName,
		Type: req.ResourceType,
		Args: args,
	}
	plid, err := backup.CreatePolicy(resource, int(req.Retention))
	if err != nil {
		log.Errorf("create backup policy failed, err: %v", err)
		return nil, err
	}
	log.Infof("create backup policy successed, policy id is %d", plid)

	return &pb.CreateBackupPolicyResponse{BackupPolicyID: int32(plid)}, nil
}

func (api *apiServer) StartBackup(ctx context.Context, req *pb.StartBackupRequest) (res *pb.StartBackupResponse, err error) {
	jobid, err := scheduling.ScheduleBackup(uint(req.ResourceID))
	if err != nil {
		log.Errorf("schedule backup job failed, err: %v", err)
		return nil, err
	}
	log.Infof("schedule backup job successed, jobid: %d", jobid)

	return &pb.StartBackupResponse{JobID: int32(jobid)}, nil
}

func Run() error {
	lis, err := net.Listen("tcp", config.Conf.Core.ListenAddress)
	if err != nil {
		log.Errorf("run server failed, err: %v", err)
		return err
	}

	log.Infof("listen at %s", config.Conf.Core.ListenAddress)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(AuthInterceptor),
		)),
	)
	reflection.Register(grpcServer)
	pb.RegisterApiServiceServer(grpcServer, &apiServer{})

	return grpcServer.Serve(lis)
}
