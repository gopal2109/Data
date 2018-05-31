package handler

import (
	"encoding/json"
	"net/http"
	"rpis/api/schema"

	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

// ThresholdsHandlerGET handles GET Request for Thresholds
func ThresholdsHandlerGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := schema.Get(schema.Thresholds{}, r.URL.Query(), 0, 0)
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

// SingleThresholdsHandlerGET gets one threshold with the offeringId from URL
func SingleThresholdsHandlerGET(w http.ResponseWriter, r *http.Request, qs httprouter.Params) {
	offeringId := qs.ByName("offeringId")
	i, err := strconv.Atoi(offeringId)
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}
	data, err := schema.GetOne(schema.Thresholds{}, "offering.offeringId", i)
	if err != nil || data == nil {
		Error(w, "Not found", http.StatusNotFound)
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

// UpdateThreshold will update specific fields allowed.
func UpdateThreshold(w http.ResponseWriter, r *http.Request, qs httprouter.Params) {
	var t schema.Thresholds
	offeringId := qs.ByName("offeringId")
	i, err := strconv.Atoi(offeringId)
	if err != nil {
		log.WithError(err).Error("not a valid offeringId")
		Error(w, "not a valid offeringId", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}

	errs := schema.Update(t, "offering.offeringId", i)
	if errs != nil {
		Error(w, errs, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	return
}
