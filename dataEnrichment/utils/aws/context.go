package aws

import (
	"context"
)

type LoggerKeyType string

var awsKey LoggerKeyType = "AWS"

func Inject(ctx context.Context, c *Connection) context.Context {
	return context.WithValue(ctx, awsKey, c)
}

func GetConnectionFromContext(ctx context.Context) *Connection {
	c, _ := ctx.Value(awsKey).(*Connection)
	return c
}
