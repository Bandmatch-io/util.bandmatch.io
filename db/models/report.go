package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ReportReasonOffensive   = "Offensive"
	ReportReasonHarrassment = "Harrassment"
	ReportReasonSpam        = "Spam"
	ReportReasonFakeProfile = "FakeProfile"
)

var ReportReasonLookup = map[string]struct{}{
	ReportReasonOffensive:   {},
	ReportReasonHarrassment: {},
	ReportReasonSpam:        {},
	ReportReasonFakeProfile: {},
}

type Report struct {
	ID                   primitive.ObjectID `bson:"_id" json:"_id"`
	Target               string             `bson:"target" json:"target"`
	ReportedUser         primitive.ObjectID `bson:"reportedUser" json:"reportedUser"`
	ReportedConversation primitive.ObjectID `bson:"reportedConversation" json:"reportedConversation"`
	Reason               string             `bson:"reason" json:"reason"`
	ExtraInformation     string             `bson:"extraInformation" json:"extraInformation"`
}

func NewReport(target primitive.ObjectID, reason string) Report {
	newRep := Report{
		ID:                   primitive.NewObjectID(),
		Target:               "User",
		ReportedUser:         target,
		ReportedConversation: primitive.NilObjectID,
		Reason:               reason,
		ExtraInformation:     "",
	}

	return newRep
}
