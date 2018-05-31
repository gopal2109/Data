package handlers

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"rpis/api/schemas"
	"rpis/api/models"
	"labix.org/v2/mgo/bson" 
	"github.com/julienschmidt/httprouter"
)

func ThresholdsHandlerGET(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	thresholds := models.GetThresholds(bson.M{})
	
	thresholdsOut := make([]schemas.Thresholds, len(thresholds))
	
	for i, t := range thresholds {
		dt := make(schemas.DatacenterThresholds)
		for k, v := range t.DatacenterThresholds {
			dt[k] = schemas.ThresholdState {
				Warning: v.Warning,
				Critical: v.Critical,
			}
		}

		thresholdsOut[i] = schemas.Thresholds{
			Offering: schemas.Offering{
				Href: t.Offering.Href,
				OfferingId: t.Offering.OfferingId,
			},
			DatacenterThresholds: dt,
		}
			
	}

	buf, err := json.Marshal(thresholdsOut)
	if err != nil {
		log.Error(fmt.Sprintf("Json Marshaling Error: %s", err.Error()))
	}
	
	Response(rw, http.StatusOK, string(buf))
}

func SingleThresholdsHandlerGET(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	offeringId := p.ByName("offeringId")
	id, err := strconv.Atoi(offeringId)

	_, threshold := models.GetThreshold(bson.M{"offering.offeringId": id}) //handle error

	//TODO: Handle Error for 404
	
	// if threshold == (models.Thresholds{}) {
	// 	Response(rw, http.StatusNotFound, "")
	// 	return
	// }
	
	dt := make(schemas.DatacenterThresholds, 0)
	for k, v := range threshold.DatacenterThresholds {
		dt[k] = schemas.ThresholdState{
			DatacenterId: v.DatacenterId,
			DatacenterAbbreviation: v.DatacenterAbbreviation,
			Warning: v.Warning,
			Critical: v.Critical,
		}
	}
	
	thresholdOut := schemas.Thresholds {
		Offering: schemas.Offering {
			Href: threshold.Offering.Href,
			OfferingId: threshold.Offering.OfferingId,
		},
		DatacenterThresholds: dt,
	}
		
	buf, err := json.Marshal(thresholdOut)
	if err != nil {
		log.Error(fmt.Sprintf("Json Marshaling Error: %s", err.Error()))
		Response(rw, http.StatusInternalServerError, "")
		return
	}

	Response(rw, http.StatusOK, string(buf))
}

func ThresholdsHandlerPUT(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := *(new(schemas.ThresholdsUpdate))
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Error(fmt.Sprintf("Json Marshaling Error: %s", err.Error()))
		Response(rw, http.StatusBadRequest, "Not Found")
		return
	}

	offeringId := p.ByName("offeringId")
	id, err := strconv.Atoi(offeringId)
	if err != nil {
		Response(rw, http.StatusNotFound, "")
		return
	}

	err, threshold := models.GetThreshold(bson.M{"offering.offeringId": id})
	if err != nil {
		Response(rw, http.StatusNotFound, "Not Found")
		return
	}	

	if t.Offering.Href != "" {
		threshold.Offering.Href = t.Offering.Href
	}

	if t.Offering.OfferingId != 0 {
		threshold.Offering.OfferingId = t.Offering.OfferingId
	}

	for _, ts := range t.DatacenterThresholds {
		threshold.DatacenterThresholds[ts.DatacenterAbbreviation] = models.ThresholdState {
			DatacenterId: ts.DatacenterId,
			DatacenterAbbreviation: ts.DatacenterAbbreviation,
			Warning: ts.Warning,
			Critical: ts.Critical,
		}
	}

	err = threshold.Update("sab")
	if err != nil {
		log.Error(err.Error())
		Response(rw, http.StatusInternalServerError, "")
		return
	}
	
	Response(rw, http.StatusAccepted, "Updated")
}
