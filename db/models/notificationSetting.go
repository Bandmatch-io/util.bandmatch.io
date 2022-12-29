package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationSetting struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id"`
	User               primitive.ObjectID `bson:"user" json:"user"`
	AllowNewsletters   bool               `bson:"allowNewsletters" json:"allowNewsletters"`
	AllowMessageAlerts bool               `bson:"allowMessageAlerts" json:"allowMessageAlerts"`
	AllowUserAlerts    bool               `bson:"allowUserAlerts" json:"allowUserAlerts"`
}

func NewNotificationSetting(user primitive.ObjectID) NotificationSetting {
	return NotificationSetting{
		ID:                 primitive.NewObjectID(),
		User:               user,
		AllowNewsletters:   true,
		AllowMessageAlerts: true,
		AllowUserAlerts:    true,
	}
}
