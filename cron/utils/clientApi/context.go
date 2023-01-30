package clientApi

import "context"

type ApiClientType string

var apiClientKey ApiClientType = "ClientApi"

func Inject(ctx context.Context, c *ClientApi) context.Context {
	return context.WithValue(ctx, apiClientKey, c)
}

func GetClientApiFromContext(ctx context.Context) *ClientApi {
	c, _ := ctx.Value(apiClientKey).(*ClientApi)
	return c
}
