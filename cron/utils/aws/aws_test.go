package aws

import (
	"context"
	"fmt"
	"github.com/Mufraggi/AWS_ETL_EXEMPLE/utils/logger"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

type mockSQS struct {
	sqsiface.SQSAPI
	messages map[string][]*sqs.Message
}

func (m *mockSQS) SendMessage(in *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	m.messages[*in.QueueUrl] = append(m.messages[*in.QueueUrl], &sqs.Message{
		Body: in.MessageBody,
	})
	return &sqs.SendMessageOutput{}, nil
}
func (m *mockSQS) ReceiveMessage(in *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	fmt.Println(m.messages[*in.QueueUrl])
	message := m.messages[*in.QueueUrl]
	return &sqs.ReceiveMessageOutput{
		Messages: message,
	}, nil
}

// Not used. This is just here for reference.

func getMockSQSClient() sqsiface.SQSAPI {
	return &mockSQS{
		messages: map[string][]*sqs.Message{},
	}
}

func TestQueue(t *testing.T) {
	log, _ := zap.NewProduction()
	ctx := context.Background()
	ctx = logger.Inject(ctx, log)
	queueUrlSqs := "https://queue.amazonaws.com/80398EXAMPLE/MyQueue"
	c := Connection{
		sqs:      getMockSQSClient(),
		queueURL: queueUrlSqs,
	}
	t.Setenv("XYZ_URL", "http://example.com")
	c.SendSqsMsg(ctx, "aaaaaa")

	message, _ := c.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &queueUrlSqs,
	})
	assert.Equal(t, *message.Messages[0].Body, "aaaaaa")
}
