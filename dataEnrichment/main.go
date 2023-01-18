package main

import (
	"context"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/tree/main/dataEnrichment/utils/aws"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/tree/main/dataEnrichment/utils/logger"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var ctx context.Context

func init() {
	log, _ := zap.NewProduction()
	c, err := aws.New()
	if err != nil {
		log.Fatal("Aws set up problems", zap.Error(err))
	}
	ctx = context.Background()
	ctx = logger.Inject(ctx, log)
	ctx = aws.Inject(ctx, c)
}

type MyEvent struct {
	Id string `json:"id"`
}

type Test struct {
	Hello string `json:"hello"`
}

func main() {
	lambda.StartWithOptions(Handler, lambda.WithContext(ctx))
}
func Handler(ctx context.Context, event interface{}) {
	log := logger.GetCLoggerFromContext(ctx)
	log.Info("receive event lambda", zap.Any("event", event))
	aws := aws.GetConnectionFromContext(ctx)
	err := aws.SendSqsMsg(ctx, "helllo sqs")
	if err != nil {
		log.Error("Couldn't send sqs msg", zap.Error(err))
	}

}
