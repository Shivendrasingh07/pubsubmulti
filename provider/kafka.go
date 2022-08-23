package provider

import (
	"context"
	"example.com/m/models"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, message string, topic models.Topic)
}

type KafkaReader interface {
	ReadMessage(ctx context.Context, groupID string)
}
