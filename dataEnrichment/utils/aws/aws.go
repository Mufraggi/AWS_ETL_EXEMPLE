package aws

import (
	"context"
	"fmt"
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

	// cfg, err := config.LoadDefaultConfig(context.Background())
	//if err != nil {
	//	return nil, err
	//}
	return &Connection{
		sqs:      sqs.New(s),
		queueURL: os.Getenv("QUEUE_URL"),
	}, nil
}

func (c *Connection) SendSqsMsg(ctx context.Context, message string) error {
	log := logger.GetCLoggerFromContext(ctx)
	fmt.Println(c.queueURL)
	messageInput := &sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    &c.queueURL,
	}

	o, err := c.sqs.SendMessage(messageInput)
	log.Info("sent message", zap.Any("output", o))
	return err
}
