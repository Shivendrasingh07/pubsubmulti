package main

import (
	"context"
	"example.com/m/models"
	"example.com/m/provider"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Message struct {
	ToUserID string
	Message  string
}

type kafkaWriter struct {
	ChatWriterChat *kafka.Writer
}

type kafkaReader struct {
	ChatReaderChat *kafka.Reader
}

func main() {

	fmt.Println("start main")
	ctx := context.Background()

	kw := publish()
	kr := subscribe(models.TopicChatMessage)
	kr2 := subscribe2(models.TopicChatMessage)

	go kr.ReadMessage(ctx, string(models.GroupChat1))
	go kr2.ReadMessage(ctx, string(models.GroupChat2))

	message := ""
	fmt.Println("Enter message")
	_, err := fmt.Scan(&message)
	if err != nil {
		panic(err)
		return
	}

	kw.WriteMessage(ctx, message, models.TopicChatMessage)
	time.Sleep(10 * time.Second)

}

func publish() provider.KafkaWriter {

	topic2 := models.TopicChatMessage

	kafkaHost := "127.0.0.1:9092"

	chatWriter := &kafka.Writer{
		Addr:      kafka.TCP(kafkaHost),
		Topic:     string(topic2),
		BatchSize: 1,
	}
	return &kafkaWriter{
		ChatWriterChat: chatWriter,
	}

}

func (k *kafkaWriter) WriteMessage(ctx context.Context, message string, topic2 models.Topic) {

	err := k.ChatWriterChat.WriteMessages(ctx,
		kafka.Message{

			Value: []byte(message),
		},
	)
	if err != nil {
		logrus.Errorf("Publish: failed to write chat messages: %v", err)
	}

	if err = k.ChatWriterChat.Close(); err != nil {
		logrus.Errorf("Publish: failed to chatWriter.Close() : %v", err)
	}
}

func subscribe(topic models.Topic) provider.KafkaReader {
	kafkaHost := "127.0.0.1:9092"
	groupID := string(models.GroupChat1)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaHost},
		Topic:   string(topic),
		GroupID: groupID,
	})
	return &kafkaReader{
		ChatReaderChat: r,
	}

}

func subscribe2(topic models.Topic) provider.KafkaReader {
	kafkaHost := "127.0.0.1:9092"
	groupID := string(models.GroupChat2)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaHost},
		Topic:   string(topic),
		GroupID: groupID,
	})
	return &kafkaReader{
		ChatReaderChat: r,
	}
}

func (kr *kafkaReader) ReadMessage(ctx context.Context, groupID string) {
	run := true
	for run {
		m, err := kr.ChatReaderChat.ReadMessage(ctx)
		if err != nil {
			logrus.Errorf("Failed to Read Message from topic: %q error: %v", m.Topic, err)
		}
		fmt.Println(groupID)
		fmt.Println(string(m.Value))
	}
}
