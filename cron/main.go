package main

import (
	"context"
	"fmt"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/utils/aws"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/utils/clientApi"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/utils/logger"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var ctx context.Context

func init() {
	log, _ := zap.NewProduction()
	c, err := aws.New()
	// todo reprendre le intialisation du aws et client api
	if err != nil {
		log.Fatal("Aws set up problems", zap.Error(err))
	}
	cApi := clientApi.New()
	ctx = context.Background()

	ctx = aws.Inject(ctx, c)
	ctx = logger.Inject(ctx, log)
	ctx = clientApi.Inject(ctx, cApi)
}

type MyEvent struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log := logger.GetCLoggerFromContext(ctx)
	c := clientApi.GetClientApiFromContext(ctx)
	a, err := c.GetListing()
	if err != nil {
		log.Error("Couldn't send sqs msg", zap.Error(err))
		return "", err
	}
	printList(a)
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func printList(list []clientApi.ArrivalRes) {
	for _, a := range list {
		fmt.Println(a)
	}
}
