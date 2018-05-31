package schema

import (
	"time"

	"rpis/backends"

	"rpis/config"

	log "github.com/Sirupsen/logrus"
	"labix.org/v2/mgo/bson"
)

// Provider ...
type Provider struct {
	Href string `json:"href" bson:"href" validate:"string,min=1,max=0"`
	ID   int    `json:"id" bson:"id" validate:"number,min=1,max=0"`
	Name string `json:"name" bson:"name" validate:"string,min=1,max=0"` // string or enum?
}

//Location structure
type Location struct {
	AggrZone             string `json:"aggrZone" bson:"aggrZone" validate:"string,min=1,max=0"`
	DatacenterID         int    `json:"datacenterId" bson:"datacenterId" validate:"number,min=1,max=0"`
	Datacenter           string `json:"datacenter" bson:"datacenter" validate:"string,min=1,max=0"`
	Cabinet              string `json:"cabinetName" bson:"cabinetName" validate:"string,min=1,max=0"`
	CabinetStartingSpace int    `json:"cabinetStartingSpace" bson:"cabinetStartingSpace" validate:"number,min=0,max=0"`
}

// ProductCatalogDetails has catalog info
type ProductCatalogDetails struct {
	Href                string `json:"href" bson:"href" validate:"string,min=1,max=0"`
	ID                  int    `json:"id" bson:"id" validate:"number,min=1,max=0"`
	OfferingDescription string `json:"offeringDescription" bson:"offeringDescription" validate:"string,min=1,max=0"`
}

// OfferServiceDetails has offer info
type OfferServiceDetails struct {
	Href        string `json:"href" bson:"href" validate:"string,min=1,max=0"`
	ProductID   string `json:"productId" bson:"productId" validate:"string,min=1,max=0"`
	ProductName string `json:"productName" bson:"productName" validate:"string,min=1,max=0"`
}

//Product structure
type Product struct {
	ProductCatalogDetails `json:"productCatalogDetails" bson:"productCatalogDetails"`
	OfferServiceDetails   `json:"offerServiceDetails" bson:"offerServiceDetails"`
}

//DeviceService has ...
type DeviceService struct {
	Href string `json:"href" bson:"href" validate:"string,min=1"`
	ID   string `json:"id" bson:"id" validate:"string,min=1"`
}

// Core ...
type Core struct {
	Href string `json:"href" bson:"href" validate:"string,min=1"`
	ID   string `json:"id" bson:"id" validate:"string,min=1"`
}

// InventoryDevice ...
type InventoryDevice struct {
	DeviceService `json:"deviceService" bson:"deviceService"`
	Core          `json:"core" bson:"core"`
}

// Links ...
type Links struct {
	Account string `json:"account" bson:"account" validate:"string,min=1"`
	Device  string `json:"device" bson:"device" validate:"string,min=1"`
}

// Account ...
type Account struct {
	Links     `json:"links" bson:"links"`
	AccountID int `json:"accountId" bson:"accountId" validate:"number,min=1"`
}

// Quote ...
type Quote struct {
	Href              string `json:"href" bson:"href" validate:"string,min=1"`
	ID                string `json:"id" bson:"id" validate:"string,min=1"`
	SalesPersonUserID string `json:"salesPersonUserId" bson:"salesPersonUserId" validate:"string,min=1"`
	Opportunity       string `json:"opportunity" bson:"opportunity" validate:"string,min=1"`
}

// AutomationEvent ...
type AutomationEvent struct {
	Href    string `json:"href" bson:"href" validate:"string,min=1"`
	EventID string `json:"eventId" bson:"eventId" validate:"string,min=1,max=0"`
}

//InventoryState structure
type InventoryState struct {
	Ref             bson.ObjectId   `bson:"ref,omitempty" validate:"-"`
	State           string          `json:"state" bson:"state" validate:"enum,min=1,AVAILABLE,MAINTENANCE,SUSPENDED,DECOMMISSIONED,ALLOCATED,NEW"`
	UserID          string          `json:"userId,omitempty" bson:"userId" validate:"string,min=1"`
	Comment         string          `json:"comment,omitempty" bson:"comment" validate:"string,min=1"`
	TimeStamp       time.Time       `json:"timestamp,omitempty" bson:"timestamp" validate:"-"`
	Device          InventoryDevice `json:"device,omitempty" bson:"device,omitempty"`
	Account         Account         `json:"account" bson:"account,omitempty"`
	Quote           Quote           `json:"quote" bson:"quote,omitempty"`
	AutomationEvent AutomationEvent `json:"automationEvent" bson:"automationEvent,omitempty"`
}

//DeviceState Structure
type DeviceState struct {
	Ref       bson.ObjectId `json:"-" bson:"ref,omitempty" validate:"-"`
	State     string        `json:"state" bson:"state" validate:"enum,min=1,TEST,PRE-PRODUCTION,PRODUCTION"`
	Comment   string        `json:"comment" bson:"comment" validate:"string,min=1,max=0"`
	TimeStamp time.Time     `json:"timestamp,omitempty" bson:"timestamp" validate:"-"`
}

func (ds *DeviceState) save(ref bson.ObjectId) error {
	session, err := backends.GetSession()
	if err != nil {
		return err
	}
	defer session.Close()

	ds.Ref = ref
	ds.TimeStamp = time.Now()
	collection := session.DB("").C(config.DeviceStateCollection)
	if err != nil {
		return err
	}

	err = collection.Insert(&ds)
	if err != nil {
		log.WithError(err).Warn("failed to save latest deviceState")
	}
	return err
}

// ChangeLog ...
type ChangeLog struct {
	UserID    string    `json:"userId" bson:"userId"`
	TimeStamp time.Time `json:"timestamp" bson:"timestamp" validate:"-"`
}

// Device schema
type Device struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty" validate:"-"`
	DeviceType     string        `json:"type" bson:"type,omitempty" validate:"enum,min=1,SERVER,FIREWALL"`
	MacAddress     string        `json:"macAddress" bson:"macAddress" validate:"string,min=1,max=0"`
	Comment        string        `json:"comment" bson:"comment,omitempty" validate:"string,min=1,max=0"`
	Deleted        bool          `json:"deleted" bson:"deleted,omitempty"`
	Location       `json:"location" bson:"location,omitempty"`
	Product        `json:"product" bson:"product,omitempty"`
	Provider       `json:"provider" bson:"provider,omitempty"`
	InventoryState `json:"inventoryState" bson:"inventoryState,omitempty" validate:"-"`
	DeviceState    `json:"deviceState" bson:"deviceState,omitempty"`
	//TODO: add dynamically the userId & create/lastupdated timestamp
}

// Save the device instance into the store
func (d Device) Save() (string, error) {
	session, err := backends.GetSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	d.ID = bson.NewObjectId()
	collection := session.DB("").C(config.DeviceCollection)
	_, err = collection.UpsertId(d.ID, &d)
	if err == nil {
		d.DeviceState.save(d.ID)
	} else {
		return "", err
	}

	log.WithFields(log.Fields{
		"id":     d.ID.Hex(),
		"action": "create",
	}).Info("device is created")
	return d.ID.Hex(), nil
}

// Find looks up for a device with filters
func (d Device) Find(filters interface{}, offset, limit int) (interface{}, error) {
	session, err := backends.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	collection := session.DB("").C(config.DeviceCollection)
	query := collection.Find(filters)
	maxResult := config.C().Application.MaxResults
	query.Skip(offset)
	if limit != 0 {
		query = query.Limit(limit)
	} else if maxResult > 0 {
		query = query.Limit(maxResult)
	} else {
		log.WithFields(log.Fields{
			"config.Application.MaxResult": maxResult,
			"default":                      100,
		}).Warn("Application.MaxResult not set, fallback to default")
		query = query.Limit(100)
	}

	if limit == 1 {
		var result Device
		err := query.One(&result)
		return result, err
	}
	//iter := query.Iter()
	var results []Device
	err = query.All(&results)
	return results, err
}

// Delete a device with ID
func (d Device) Delete(key string, value interface{}) []error {
	u := make(map[string]interface{}, 1)
	u["$set"] = struct {
		Deleted           bool   `bson:"deleted"`
		Comment           string `bson:"comment" validate:"string,min=1,max=0"`
		AutomationEventID string `bson:"inventoryState.automationEvent.eventId" validate:"string,min=1,max=0"`
	}{
		true,
		d.Comment,
		d.AutomationEvent.EventID,
	}

	errs := ValidateFields("device", u["$set"])
	if len(errs) > 0 {
		return errs
	}

	selector := map[string]interface{}{key: value}
	session, err := backends.GetSession()
	if err != nil {
		return []error{err}
	}
	defer session.Close()

	collection := session.DB("").C(config.DeviceCollection)
	if err = collection.Update(selector, &u); err != nil {
		log.WithError(err).WithFields(log.Fields{
			"key":   key,
			"value": value,
		}).Error("Failed to delete device")
		return []error{err}
	}
	log.WithFields(log.Fields{
		"key":    key,
		"value":  value,
		"action": "delete",
	}).Info("device is deleted")
	return nil
}

// Update an device with all omitempty ignored.
func (d Device) Update(key string, value interface{}) []error {
	u := make(map[string]interface{}, 1)
	u[key] = value

	// required to restrict other fields from being modified
	toUpdate := struct {
		DeviceType string `bson:"type" validate:"enum,min=1,SERVER,FIREWALL"`
		MacAddress string `bson:"macAddress" validate:"string,min=1,max=0"`
		Comment    string `bson:"comment" validate:"string,min=1,max=0"`
	}{
		d.DeviceType,
		d.MacAddress,
		d.Comment,
	}
	errs := ValidateFields("device", toUpdate)
	if len(errs) > 0 {
		return errs
	}

	t := map[string]interface{}{"$set": d}
	session, err := backends.GetSession()
	if err != nil {
		return []error{err}
	}
	defer session.Close()

	collection := session.DB("").C(config.DeviceCollection)
	if err = collection.Update(u, &t); err != nil {
		return []error{err}
	}
	log.WithFields(log.Fields{
		"key":    key,
		"value":  value,
		"action": "update",
	}).Info("device is updated")
	return nil
}
