package models

import (
	"time"
	"rpis/backend"
	"labix.org/v2/mgo/bson"
)

var log = backend.Log

type ThresholdState struct {
	DatacenterId int `bson:"datacenterId,omitempty"`
	DatacenterAbbreviation string `bson:"datacenterAbbreviation,omitempty"`
	Warning int `bson:"warning,omitempty"`
	Critical int `bson:"critical,omitempty"`
}

type DatacenterThresholds map[string]ThresholdState

type Thresholds struct {
	Id bson.ObjectId `bson:"_id"`
	Offering Offering `bson:"offering,omitempty"`
	DatacenterThresholds DatacenterThresholds `bson:"datacenterThresholds,omitempty"`
	LastModified ChangeLog `bson:"lastModified"`
}

type Offering struct {
	Href string `bson:"href,omitempty"`
	OfferingId int `bson:"offeringId,omitempty"`
}

func GetThresholds(filter bson.M) []Thresholds {
	var t []Thresholds
	err := backend.GetDB().C("Thresholds").Find(filter).All(&t)
	if err != nil {
		log.Error(err.Error())
	}
	return t
}

func GetThreshold(filter bson.M) (error, Thresholds) {
	var t Thresholds
	err := backend.GetDB().C("Thresholds").Find(filter).One(&t)
	return err, t
}

func (t Thresholds) Update(userId string) error {
	if t.Id.String() != "" {
		t.LastModified.TimeStamp = time.Now()
		t.LastModified.UserId = userId
		_, err := backend.GetDB().C("Thresholds").UpsertId(t.Id, &t)
		return err
	} else {
		log.Warning("Cannot Create Threshold this way")
		return ValidationError{"Cannot create thershold in DB", 400}
	}
}
