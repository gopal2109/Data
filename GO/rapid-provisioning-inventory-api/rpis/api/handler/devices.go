package handler

import (
	"encoding/json"
	"net/http"
	"rpis/api/schema"

	"io/ioutil"

	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"labix.org/v2/mgo/bson"
)

// Devices handles GET Request for devices
func Devices(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	data, err := schema.Get(schema.Device{}, r.URL.Query(), offset, limit)
	if err != nil {
		log.WithError(err).Error("Internal Server Error")
		Error(w, err, http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(data)
	if err != nil {
		log.WithError(err).Error("Error marshalling JSON")
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

// Device gets one device with the id from URL
func Device(w http.ResponseWriter, r *http.Request, qs httprouter.Params) {
	id := qs.ByName("id")
	if !bson.IsObjectIdHex(id) {
		Error(w, "not found", http.StatusNotFound)
		return
	}
	i := bson.ObjectIdHex(id)
	data, err := schema.GetOne(schema.Device{}, "_id", i)
	if err != nil || data == nil {
		Error(w, "not found", http.StatusNotFound)
		return
	}

	buf, err := json.Marshal(data)
	if err != nil {
		log.WithError(err).Error("Error marshalling JSON")
		Error(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

// CreateDevice is to create a new device in the collection
func CreateDevice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var devices []schema.Device
	var buf []byte
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(buf, &devices)
	if err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}
	dModel := make([]schema.Model, len(devices))
	for i := range devices {
		dModel[i] = devices[i]
	}
	ids, errs := schema.Create(dModel...) // TODO errs
	if len(errs) > 0 {
		Error(w, errs, http.StatusTeapot)
		return
	}

	res := make([]interface{}, 0)
	for _, id := range ids {
		respData := struct {
			Href string `json:"href"`
			ID   string `json:"id"`
		}{
			"../inventory/devices/" + id,
			id,
		}
		res = append(res, respData)
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.WithError(err).Error("failed to encode error message")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

// DeleteDevice delete device by key
func DeleteDevice(w http.ResponseWriter, r *http.Request, qs httprouter.Params) {
	id := qs.ByName("id")
	if !bson.IsObjectIdHex(id) {
		Error(w, "not found", http.StatusNotFound)
		return
	}
	i := bson.ObjectIdHex(id)

	var buf []byte
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}

	reqPayload := struct {
		Comment           string `json:"comment"`
		AutomationEventID string `json:"automationEventId"`
	}{}

	err = json.Unmarshal(buf, &reqPayload)
	if err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}
	d := schema.Device{
		Comment: reqPayload.Comment,
	}
	d.InventoryState.AutomationEvent.EventID = reqPayload.AutomationEventID

	// Check if resource already deleted
	x, errs := schema.GetOne(d, "_id", i)
	if errs != nil {
		Error(w, errs, http.StatusNotFound)
		return
	} else if x.(schema.Device).Deleted == true {
		Error(w, "you are trying a delete a device already deleted", http.StatusGone)
		return
	}
	if errs := schema.Delete(d, "_id", i); len(errs) > 0 {
		Error(w, errs, http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// UpdateDevice will update specific fields allowed.
func UpdateDevice(w http.ResponseWriter, r *http.Request, qs httprouter.Params) {
	var d schema.Device
	id := qs.ByName("id")
	if !bson.IsObjectIdHex(id) {
		Error(w, "not found", http.StatusNotFound)
		return
	}
	i := bson.ObjectIdHex(id)

	var buf []byte
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(buf, &d)
	if err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}

	errs := schema.Update(d, "_id", i)
	if errs != nil {
		Error(w, errs, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	return
}
