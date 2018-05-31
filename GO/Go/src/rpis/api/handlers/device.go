package handlers

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"rpis/config"
	"encoding/json"
	"rpis/api/models"
	"rpis/api/schemas"
	"labix.org/v2/mgo/bson"
	"github.com/julienschmidt/httprouter"
	lr "github.com/Sirupsen/logrus"

)

func DeviceHandlerPOST(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	d := *(new(schemas.DeviceCreate))
	
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		Response(rw, http.StatusBadRequest, "Bad request")
		return
	}
	
	err, device := models.NewDevice(d.DeviceType, d.MacAddress,
		d.Comment, d.Deleted, d.Provider.Id,
		d.Provider.Name, d.Location.AggrZone,
		d.Location.DatacenterId, d.Location.Datacenter,
		d.Location.Cabinet, d.Location.CabinetStartingSpace,
		d.Product.ProductCatalogDetails.Id,
		d.Product.ProductCatalogDetails.OfferingDescription,
		d.Product.OfferServiceDetails.ProductId,
		d.Product.OfferServiceDetails.ProductName,
		d.DeviceState.State,
		d.DeviceState.Comment,
		"Sab")
	
	if err != nil {
		e := err.(models.ValidationError)
		Response(rw, e.Code, e.Json())
		return
	}

	err = device.Save("Sab")
	
	if err != nil {
		log.WithFields(lr.Fields{"error": err.Error()}).Info("Internal Server Error")
		Response(rw, http.StatusInternalServerError, "InternalServerError")
		return
	}
	
	Response(rw, http.StatusCreated, "Created")
	return
}

func DeviceHandlerGET(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/* query params */
	var (
		offset int = 0
		limit int = 10
	)

	if p.ByName("limit") != "" {
		l, err := strconv.Atoi(p.ByName("limit"))
		if err == nil {
			limit = l
		}
	}
	if p.ByName("offset") != "" {
		o, err := strconv.Atoi(p.ByName("offset"))
		if err == nil {
			offset = o
		}
	}

	urlQueries := r.URL.Query()
	delete(urlQueries, "limit")
	delete(urlQueries, "offset")

	// searchable fields
	
	filter := make(bson.M, 0)
	
	for q, v := range urlQueries {
		val := models.DeviceSearchableFields[q]
		switch val {
		case "string":
			filter[q] = v[0]
		case "bool":
			x, err := strconv.ParseBool(v[0])
			if err == nil {
				filter[q] = x
			}
		case "int":
			x, err := strconv.ParseInt(v[0], 10, 16)
				if err == nil {
					filter[q] = x
				}
		case "timestamp":
			loc, err := time.LoadLocation(config.Conf.Application.TimeZone)
			if err != nil {
				log.Warn("Config[Application.TimeZone] error, assuming UTC")
					loc, _ = time.LoadLocation("UTC")
			}
				t, err := time.ParseInLocation(config.Conf.Application.DateFormat, v[0], loc)
			if err != nil {
					log.Warn(fmt.Sprintf("Skipping search param \"%s\"", q))
				break
			} else {
				filter[q] = bson.M{"$gte": t}
			}
		default:
			log.WithFields(lr.Fields{"field": q}).Info(fmt.Sprintf("Ignoring non searchable parameter %s", q))
		}
	}
	
	res := models.GetDevices(filter, offset, limit)
	buf, _ := json.Marshal(res)
	Response(rw, http.StatusOK, string(buf))
}
	
func SingleDeviceHandlerGET(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	deviceId := p.ByName("deviceId")
	err, device := models.GetOneDevice(deviceId)
	if err != nil {
		Response(rw, http.StatusNotFound, "Not Found")
		return
	} 
	buf, err := json.Marshal(device)
	if err != nil {
		/* TODO: Write common error handling methods */
	}
	Response(rw, http.StatusOK, string(buf))
}

func SingleDeviceHandlerPUT(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	d := *(new(schemas.DeviceCreate))
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		Response(rw, http.StatusBadRequest, "Not Found")
		return
	}

	deviceId := p.ByName("deviceId")
	
	err, device := models.GetOneDevice(deviceId)
	if err != nil {
		Response(rw, http.StatusNotFound, "Not Found")
		return
	}
	
	if d.DeviceType != "" {
		device.DeviceType = d.DeviceType
	}
	if d.MacAddress != "" {
		device.MacAddress = d.MacAddress
	}
	if d.Comment != "" {
		device.Comment = d.Comment
	}

	err = device.Save("sab")

	if err != nil {
		Response(rw, http.StatusBadRequest, string(err.Error()))
		return
	}

	Response(rw, http.StatusAccepted, "Updated")
}

func SingleDeviceHandlerDELETE(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	deviceId := p.ByName("deviceId")
	err, device := models.GetOneDevice(deviceId)
	if err != nil {
		Response(rw, http.StatusNotFound, "Bad request")
		return
	}

	device.Deleted = true
	err = device.Save("sab")
	if err != nil {
		Response(rw, http.StatusNotFound, "Not found")
		return
	}
	
	Response(rw, http.StatusNoContent, "")
}

func InventoryStatesHandlerPOST(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	deviceId := p.ByName("deviceId")
	err, device := models.GetOneDevice(deviceId)
	if err != nil {
		Response(rw, http.StatusNotFound, "Not found")
		return
	}
	
	d := *(new(schemas.InventoryState)) 
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		Response(rw, http.StatusBadRequest, "Bad request")
		return
	}
	
	device.InventoryState.Id = bson.NewObjectId()
	device.InventoryState.Device = device.Id
	device.InventoryState.UserId = "Sab"
	device.InventoryState.TimeStamp = time.Now()
	
	if d.State != "" {
		err := device.InventoryState.SetState(d.State)
		if err != nil {
			Response(rw, http.StatusBadRequest, string(err.Error()))
			return
		}
	}

	if d.Comment != "" {
		device.InventoryState.Comment = d.Comment
	} else {
		Response(rw, http.StatusBadRequest, "Comment Required")
		return
	}

	if d.Account.AccountId != "" {
		device.InventoryState.Account.AccountId = d.Account.AccountId
	}

	if d.AutomationEvent.EventId != "" {
		device.InventoryState.AutomationEvent.EventId = d.AutomationEvent.EventId
	}

	if d.Quote.Id != "" {
		device.InventoryState.Quote.Id = d.Quote.Id
	}

	if d.Quote.SalesPersonUserId != "" {
		device.InventoryState.Quote.SalesPersonUserId = d.Quote.SalesPersonUserId
	}

	if d.Quote.Opportunity != "" {
		device.InventoryState.Quote.Opportunity = d.Quote.Opportunity
	}

	err = device.Save("Sab")
	if err != nil {
		Response(rw, 422, err.Error())
		return
	}

	device.InventoryState.Save()
	if err != nil {
		Response(rw, 500, "backup failed")
		return
	}

	Response(rw, http.StatusAccepted, "Updated")
}
// - Look up for the device
// - request body, update the InventoryState struct
// - Save in the device document
// - save the inventorystate struct in the IS collection

func InventoryStatesHandlerGET(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	deviceId := p.ByName("deviceId")
	err, device := models.GetOneDevice(deviceId)
	if err != nil {
		Response(rw, http.StatusNotFound, "Not found")
		return
	}

	inventoryStates := models.GetInventoryStates(device.Id)
	var res []schemas.InventoryState
	
	for _, i := range inventoryStates {
		x := schemas.InventoryState{
			State: i.State,
			Comment: i.Comment,
			AutomationEvent: i.AutomationEvent,
			Account: i.Account,
			Quote: i.Quote}
		res = append(res, x)
	}
	
	buf, err := json.Marshal(res)
	if err != nil {
		Response(rw, http.StatusInternalServerError, "")
		return
	}
	
	Response(rw, http.StatusOK, string(buf))
}
