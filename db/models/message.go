package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Conversation primitive.ObjectID `bson:"conversation" json:"conversation"`
	Content      string             `bson:"content" json:"content"`
	Sender       primitive.ObjectID `bson:"sender" json:"sender"`
	Read         bool               `bson:"read" json:"read"`
	Timestamp    time.Time          `bson:"timestamp" json:"timestamp"`
}

func NewMessage(conversation primitive.ObjectID, content string, sender primitive.ObjectID) Message {
	newMsg := Message{
		ID:           primitive.NewObjectID(),
		Conversation: conversation,
		Content:      content,
		Read:         false,
		Sender:       sender,
		Timestamp:    time.Now(),
	}

	return newMsg
}
