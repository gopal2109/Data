package models

import (
	"fmt"
	"time"
	"encoding/json"
)

const (
	ErrorUnprocessable = 422
	/* TODO: Error codes according to kind of errors */
)

type ChangeLog struct {
	UserId string `json:"userId" bson:"userId"`
	Comment string `json:"comment" bson:"comment"`
	TimeStamp time.Time `json:"timestamp" bson:"timestamp"`
}

type ValidationError struct {
	What string `json:"error"`
	Code int `json:"code"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.What)
}

func (e ValidationError) Json() string {
	s, _ := json.Marshal(e)
	return string(s)
}
