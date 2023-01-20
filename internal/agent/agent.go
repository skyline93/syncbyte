package agent

import (
	"context"
	"errors"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/skyline93/syncbyte-go/internal/proto"
)

type syncbyteServer struct {
	pb.UnimplementedSyncbyteServer
}

func (s *syncbyteServer) Backup(req *pb.BackupRequest, stream pb.Syncbyte_BackupServer) error {
	ctx := context.TODO()

	var stor DataStore

	switch Conf.Storage.Type {
	case NAS:
		stor = NewNasVolume()
	// case S3:
	default:
		return errors.New("storage type not support")
	}

	mgmt := NewBackupManager(stor, ctx)

	fiChan := make(chan FileInfo)

	go mgmt.Backup(req.SourcePath, fiChan)

	for fi := range fiChan {
		var infos []*pb.PartInfo

		for _, i := range fi.PartInfos {
			info := &pb.PartInfo{
				Index: int32(i.Index),
				MD5:   i.MD5,
				Size:  i.Size,
			}

			infos = append(infos, info)
		}

		if err := stream.Send(&pb.BackupResponse{
			Name:       fi.Name,
			Path:       fi.Path,
			Size:       fi.Size,
			MD5:        fi.MD5,
			GID:        fi.GID,
			UID:        fi.UID,
			Device:     fi.Device,
			DeviceID:   fi.DeviceID,
			BlockSize:  fi.BlockSize,
			Blocks:     fi.Blocks,
			AccessTime: fi.AccessTime,
			ModTime:    fi.ModTime,
			ChangeTime: fi.ChangeTime,
			PartInfos:  infos,
		}); err != nil {
			return err
		}
	}

	return nil
}

func newServer() *syncbyteServer {
	return &syncbyteServer{}
}

func RunServer() error {
	lis, err := net.Listen("tcp", Conf.Core.GrpcAddr)
	if err != nil {
		log.Errorf("run server failed, err: %v", err)
		return err
	}

	log.Infof("listen at %s", Conf.Core.GrpcAddr)

	grpcServer := grpc.NewServer()
	pb.RegisterSyncbyteServer(grpcServer, newServer())

	return grpcServer.Serve(lis)
}
