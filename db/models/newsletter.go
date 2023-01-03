package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Newsletter struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Delivered  bool               `bson:"delivered" json:"delivered"`
	Content    string             `bson:"content" json:"content"`
	Markdown   string             `bson:"markdown" json:"markdown"`
	Title      string             `bson:"title" json:"title"`
	Timestamps struct {
		Created      time.Time `bson:"created" json:"created"`
		Edited       time.Time `bson:"edited" json:"edited"`
		DeliveryTime time.Time `bson:"deliveryTime" json:"deliveryTime"`
	} `bson:"timestamps" json:"timestamps"`
}

func NewNewsletter(content string, markdown string, title string, deliveryTime time.Time) Newsletter {
	nl := Newsletter{
		ID:        primitive.NewObjectID(),
		Delivered: false,
		Content:   content,
		Markdown:  markdown,
		Title:     title,
	}

	nl.Timestamps.Created = time.Now()
	nl.Timestamps.DeliveryTime = deliveryTime

	return nl
}
