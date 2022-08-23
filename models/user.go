package models

import (
	"mime/multipart"
)

type Chunk struct {
	UploadID      string // unique id for the current upload.
	ChunkNumber   int32
	TotalChunks   int32
	TotalFileSize int64 // in bytes
	Filename      string
	Data          *multipart.Part
	UploadDir     string
	ByteData      []byte
}

type Message struct {
	ToUserID string
	Message  string
}

type Topic string

const (
	TopicChatMessage     Topic = "chat_message_new4"
	TopicRealtimeMessage Topic = "real_time_events_message"
)

type Group string

const (
	GroupChat1 Group = "g21"
	GroupChat2 Group = "g22"
)
