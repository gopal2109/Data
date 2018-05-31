package schemas

import (
	"rpis/api/models"
)

type DeviceState struct {
	State string `json:"state"`
	Comment string `json:"comment"`
}

type DeviceCreate struct {
	DeviceType string `json:"type"`
	MacAddress string `json:"macAddress"`
	Comment string `json:"comment"`
	Deleted bool `json:"deleted"`
	Provider models.Provider `json:"provider"`
	Location models.Location `json:"location"`
	Product models.Product `json:"product"`
	DeviceState DeviceState `json:"deviceState"`
	Created models.ChangeLog `json:"created"`
	LastModified models.ChangeLog `json:"lastModified"`
}
