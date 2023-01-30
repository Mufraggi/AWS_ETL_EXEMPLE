package aws

import (
	"context"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/tree/main/dataEnrichment/utils/logger"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"go.uber.org/zap"
	"os"
)

type Connection struct {
	sqs      sqsiface.SQSAPI
	queueURL string
}

func New() (*Connection, error) {
	s := session.Must(session.NewSession())
	return &Connection{
		sqs:      sqs.New(s),
		queueURL: os.Getenv("QUEUE_URL"),
	}, nil
}

func (c *Connection) SendSqsMsg(ctx context.Context, message string) error {
	log := logger.GetCLoggerFromContext(ctx)
	messageInput := &sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    &c.queueURL,
	}

	o, err := c.sqs.SendMessage(messageInput)
	log.Info("sent message", zap.Any("output", o))
	return err
}
