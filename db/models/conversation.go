package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID           primitive.ObjectID   `bson:"_id" json:"_id"`
	Participants []primitive.ObjectID `bson:"participants" json:"participants"`
	LastMessage  primitive.ObjectID   `bson:"lastMessage" json:"lastMessage"`
}

func NewConversation(participants []primitive.ObjectID) Conversation {
	newConvo := Conversation{
		ID:           primitive.NewObjectID(),
		Participants: participants,
		LastMessage:  primitive.NilObjectID,
	}

	return newConvo
}
