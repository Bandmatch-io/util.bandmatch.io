package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Statistic struct {
	ID           primitive.ObjectID  `bson:"_id" json:"_id"`
	Date         time.Time           `bson:"date" json:"date"`
	MessagesSent int                 `bson:"messagesSent" json:"messagesSent"`
	Logins       int                 `bson:"logins" json:"logins"`
	Signups      int                 `bson:"signups" json:"signups"`
	Searches     int                 `bson:"searches" json:"searches"`
	RootViews    int                 `bson:"rootViews" json:"rootViews"`
	Referrers    []Referrer          `bson:"referrers" json:"referrers"`
	Reports      int                 `bson:"reports" json:"reports"`
	ServerErrors int                 `bson:"serverErrors" json:"serverErrors"`
	UserErrors   int                 `bson:"userErrors" json:"userErrors"`
	MatchCount   MatchCount          `bson:"matchCount" json:"matchCount"`
	Timing       map[string]TimeData `bson:"endpointTiming" json:"endpointTiming"`
}

type MatchCount struct {
	Average int `bson:"avg" json:"avg"`
	Total   int `bson:"total" json:"total"`
}

type Referrer struct {
	URL   string `bson:"url" json:"url"`
	Count int    `bson:"count" json:"count"`
}

type TimeData struct {
	Min   float64 `bson:"min" json:"min"`
	Avg   float64 `bson:"avg" json:"avg"`
	Max   float64 `bson:"max" json:"max"`
	Count int     `bson:"count" json:"count"`
}

func NewStatistic() Statistic {
	newStat := Statistic{
		ID:        primitive.NewObjectID(),
		Date:      time.Now(),
		Timing:    map[string]TimeData{},
		Referrers: []Referrer{},
	}

	return newStat
}
