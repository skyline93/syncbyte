package grpc

import (
	"context"

	"github.com/skyline93/syncbyte-go/internal/engine/config"
	pb "github.com/skyline93/syncbyte-go/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	c       pb.SyncbyteClient
	context context.Context
}

func NewClient(ctx context.Context) (*Client, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(config.Conf.GrpcServerAddr, opts...)
	if err != nil {
		return nil, err
	}

	client := pb.NewSyncbyteClient(conn)

	return &Client{
		conn:    conn,
		c:       client,
		context: ctx,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
