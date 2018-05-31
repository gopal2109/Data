package handlers

import (
	//"fmt"
	"strconv"
	"net/http"
	"rpis/api/schemas"
	"rpis/api/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
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
	
	res := models.GetDevices(offset, limit)
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
		Response(rw, http.StatusNotFound, "Not request")
		return
	}

	device.Deleted = true
	err = device.Save("sab")
	if err != nil {
		Response(rw, http.StatusNotFound, "Not request")
		return
	}
	
	Response(rw, http.StatusNoContent, "")
}
