package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	ID          primitive.ObjectID `bson:"_id"`
	ServiceName string             `bson:"serviceName"`
	Port        int                `bson:"port"`
	Identifier  string             `bson:"identifier"`
	Version     int                `bson:"version"`
	RunnerArgs  []string           `bson:"runArgs"`
	BuildArgs   []string           `bson:"buildArgs"`
	ReleasePort int                `bson:"releasePort"`
	Exposed     bool               `bson:"exposed"`
}
