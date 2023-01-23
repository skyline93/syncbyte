package grpc

import (
	"context"

	"github.com/skyline93/syncbyte-go/internal/proto"
)

type NasResourceArgs struct {
	Dir string `json:"dir"`
}

func CreateBackupPolicy(resourceName, resourceType string, resourceArgs interface{}, retention int) (policyID uint, err error) {
	client, err := NewClient(context.TODO())
	if err != nil {
		return 0, err
	}
	defer client.Close()

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var backup_path string
	if resourceType == "nas" {
		args := resourceArgs.(NasResourceArgs)
		backup_path = args.Dir
	}

	resp, err := client.C.CreateBackupPolicy(ctx, &proto.CreateBackupPolicyRequest{
		ResourceName: resourceName,
		ResourceType: resourceType,
		Retention:    int32(retention),
		BackupPath:   backup_path,
	})
	if err != nil {
		return 0, err
	}

	return uint(resp.BackupPolicyID), nil
}

func StartBackup(resourceID uint) (jobID uint, err error) {
	client, err := NewClient(context.TODO())
	if err != nil {
		return 0, err
	}
	defer client.Close()

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	resp, err := client.C.StartBackup(ctx, &proto.StartBackupRequest{
		ResourceID: int32(resourceID),
	})
	if err != nil {
		return 0, err
	}

	return uint(resp.JobID), nil
}
