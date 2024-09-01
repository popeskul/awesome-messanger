package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"net"
)

type GrpcClient struct {
	conn    *grpc.ClientConn
	options []grpc.DialOption
}

func NewClient(options ...grpc.DialOption) *GrpcClient {
	return &GrpcClient{
		options: options,
	}
}

func (c *GrpcClient) Connect(address string, timeout time.Duration) error {
	options := append(c.options, grpc.WithBlock(), grpc.WithTimeout(timeout))

	conn, err := grpc.NewClient(address, options...)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *GrpcClient) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}

func (c *GrpcClient) GetConnection() *grpc.ClientConn {
	return c.conn
}

type ClientBuilder struct {
	options []grpc.DialOption
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{}
}

func (b *ClientBuilder) WithInsecure() *ClientBuilder {
	b.options = append(b.options, grpc.WithInsecure())
	return b
}

func (b *ClientBuilder) WithContextDialer(dialer func(context.Context, string) (net.Conn, error)) *ClientBuilder {
	b.options = append(b.options, grpc.WithContextDialer(dialer))
	return b
}

func (b *ClientBuilder) WithOptions(options ...grpc.DialOption) *ClientBuilder {
	b.options = append(b.options, options...)
	return b
}

func (b *ClientBuilder) Build() *GrpcClient {
	return NewClient(b.options...)
}
