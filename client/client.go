// package client содержит реализацию клиента.
package client

import (
	"context"

	"github.com/pkg/errors"
	greetingv1 "github.com/viktor-titov/grpc-gateway-example/client/internal/greeting/v1"
	"google.golang.org/grpc"
)

type Client struct {
	greetingClient greetingv1.GreeterServiceClient
}

func NewClient(connection *grpc.ClientConn) (*Client, error) {
	greetingClient := greetingv1.NewGreeterServiceClient((connection))
	return &Client{greetingClient}, nil
}

func (c *Client) SayHello(
	ctx context.Context,
	name string,
) (*greetingv1.SayHelloResponse, error) {
	request := greetingv1.SayHelloRequest{
		Name: name,
	}
	res, err := c.greetingClient.SayHello(ctx, &request)

	if err != nil {
		return nil, errors.Wrap(err, "call SayHello method")
	}

	return res, nil
}
