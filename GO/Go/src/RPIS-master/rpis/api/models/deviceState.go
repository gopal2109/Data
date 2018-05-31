package models

import (
	"time"
	"strings"
	"labix.org/v2/mgo/bson"
	"rpis/backend"
)

type DeviceState struct {
	State string `json:"state" bson:"state"`
	Ref bson.ObjectId `bson:"ref"`
	Log ChangeLog `json:"changelog" bson:",inline"`
}

func (d *DeviceState) SetState(state string) error {
	switch strings.ToUpper(state) {
	case "TEST":
		d.State = "TEST"
		return nil
	case "PRE-PRODUCTION":
		d.State = "PRE-PRODUCTION"
		return nil
	case "PRODUCTION":
		d.State = "PRODUCTION"
		return nil
	default:
		return ValidationError{"State should be one of TEST, PRE-PRODUCTION, PRODUCTION", ErrorUnprocessable}
	}
}

func NewDeviceState(state string, device bson.ObjectId, comment string, userid string) (error, DeviceState) {
	d := DeviceState{}
	if err := d.SetState(state); err != nil {
		return err, d
	}
	d.Ref = device
	d.Log = ChangeLog{UserId: userid, Comment: comment}
	return nil, d
}
	
func (d DeviceState) Save() error {
	d.Log.TimeStamp = time.Now()
	err := backend.GetDB().C("DeviceStates").Insert(&d)
	return err
}
