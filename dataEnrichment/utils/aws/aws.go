package aws

import (
	"context"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/tree/main/dataEnrichment/utils/logger"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/zap"
	"os"
)

type Connection struct {
	sqs      *sqs.Client
	queueURL string
}

func New() (*Connection, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	return &Connection{
		sqs:      sqs.NewFromConfig(cfg),
		queueURL: os.Getenv("QUEUE_URL"),
	}, nil
}

func (c *Connection) SendSqsMsg(ctx context.Context, message string) error {
	log := logger.GetCLoggerFromContext(ctx)
	o, err := c.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    &c.queueURL,
	})
	log.Info("sent message", zap.Any("output", o))
	return err
}
