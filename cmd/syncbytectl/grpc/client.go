package grpc

import (
	"context"

	"github.com/skyline93/syncbyte-go/cmd/syncbytectl/config"
	pb "github.com/skyline93/syncbyte-go/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	C       pb.ApiServiceClient
	context context.Context
}

func NewClient(ctx context.Context) (*Client, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(config.Conf.Core.ServerAddress, opts...)
	if err != nil {
		return nil, err
	}

	client := pb.NewApiServiceClient(conn)

	return &Client{
		conn:    conn,
		C:       client,
		context: ctx,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
