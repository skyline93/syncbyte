package grpc

import (
	"context"
	"io"

	"github.com/skyline93/syncbyte-go/internal/engine/schema"
	pb "github.com/skyline93/syncbyte-go/internal/proto"
)

func Backup(fiChan chan schema.FileInfo, sourcePath, mountPoint string, ctx context.Context) error {
	client, err := NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	req := pb.BackupRequest{
		BackupParams:    &pb.BackupParams{SourcePath: sourcePath},
		DataStoreParams: &pb.DataStoreParams{MountPoint: mountPoint},
	}

	stream, err := client.c.Backup(ctx, &req)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		var partInfos []*schema.PartInfo
		for _, i := range resp.PartInfos {
			partInfo := schema.PartInfo{
				Index: int(i.Index),
				MD5:   i.MD5,
				Size:  i.Size,
			}

			partInfos = append(partInfos, &partInfo)
		}

		fi := schema.FileInfo{
			Name:       resp.Name,
			Size:       resp.Size,
			GID:        resp.GID,
			UID:        resp.UID,
			Device:     resp.Device,
			DeviceID:   resp.DeviceID,
			BlockSize:  resp.BlockSize,
			Blocks:     resp.Blocks,
			AccessTime: resp.AccessTime,
			ModTime:    resp.ModTime,
			ChangeTime: resp.ChangeTime,
			Path:       resp.Path,
			MD5:        resp.MD5,
			PartInfos:  partInfos,
		}

		fiChan <- fi
	}

	close(fiChan)
	return nil
}
