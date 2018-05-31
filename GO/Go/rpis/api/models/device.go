package models

import (
	"fmt"
	"time"
	"rpis/api/backend"
	"labix.org/v2/mgo/bson" 
)

type Provider struct {
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func NewProvider(id int, name string) (error, Provider) {
	var p Provider
	if name == "" {
		return ValidationError{"Provider[name] cannot be empty", 400}, p
	}
	if id < 0 {
		return ValidationError{"Invalid Provider[id]", 400}, p
	}
	return nil, Provider{Id: id, Name: name}
}

type Location struct {
	AggrZone                string `json:"aggrZone" bson:"aggrZone"`
	DatacenterId            int `json:"datacenterId" bson:"datacenterId"`
	Datacenter              string `json:"datacenter" bson:"datacenter"`
	Cabinet                 string `json:"cabinetName" bson:"cabinetName"`
	CabinetStartingSpace    int `json:"cabinetStartingSpace" bson:"cabinetStartingSpace"`
}

func NewLocation(aggrZone string, datacenterId int, datacenter string, cabinetName string, cabinetStartingSpace int) (error, Location) {
	var location Location
	if aggrZone == "" {
		return ValidationError{"Location[AggrZone] is required", 400}, location
	}
	if datacenter == "" {
		return ValidationError{"Location[datacenterId] cannot be empty", 400}, location
	}
	if cabinetName == "" {
		return ValidationError{"Location[cabinetName] is required", 400}, location
	}
	if datacenterId < 1 {
		return ValidationError{"Invalid Location[datacenterId]", 400}, location
	}
	return nil, Location{aggrZone, datacenterId, datacenter, cabinetName, cabinetStartingSpace}
}

type ProductCatalogDetails struct {
	Id int `json:"id" bson:"id"`
	OfferingDescription string `json:"offeringDescription" bson:"offeringDescription"`
}

func NewProductCatalogDetails(id int, OfferingDescription string) (error, ProductCatalogDetails) {
	var productCatalog ProductCatalogDetails
	if OfferingDescription == "" {
		return ValidationError{"OfferingDescription cannot be empty", 400}, productCatalog
	}
	if id < 0 {
		return ValidationError{"productId cannot be empty", 400}, productCatalog
	}

	return nil, ProductCatalogDetails{id, OfferingDescription}
}

type OfferServiceDetails struct {
	ProductId string `json:"productId" bson:"productId"`
	ProductName string `json:"productName" bson:"productName"`
}

func NewOfferServiceDetails(productId string, productName string) (error, OfferServiceDetails) {
	var offer OfferServiceDetails
	if productName == "" {
		return ValidationError{"productName cannot be empty", 400}, offer
	}
	if productId == "" {
		return ValidationError{"Invalid productId", 400}, offer
	}

	return nil, OfferServiceDetails{productId, productName}
}

type Product struct {
	ProductCatalogDetails ProductCatalogDetails `json:"productCatalogDetails" bson:"productCatalogDetails"`
	OfferServiceDetails OfferServiceDetails `json:"offerServiceDetails" bson:"offerServiceDetails"`
}

func NewProduct(id int, OfferingDescription string, productId string, productName string) (error, Product) {
	product := *(new(Product))
	
	err, pcd := NewProductCatalogDetails(id, OfferingDescription)
	if err != nil {
		return err, product
	}
	product.ProductCatalogDetails = pcd
	
	err, osd := NewOfferServiceDetails(productId, productName)
	if err != nil {
		return err, product
	}
	product.OfferServiceDetails = osd

	return nil, product
}

type Device struct { /* Device Model */
	Id bson.ObjectId `bson:"_id"`
	DeviceType string `json:"type" bson:"type,omitempty"`
	MacAddress string `json:"macAddress" bson:"macAddress,omitempty"`
	Comment string `json:"comment" bson:"comment,omitempty"`
	Deleted bool `json:"deleted" bson:"deleted,omitempty"`
	Provider Provider `json:"provider" bson:"provider,omitempty"`
	Location Location `json:"location" bson:"location,omitempty"`
	Product Product `json:"product" bson:"product,omitempty"`
	InventoryState `json:"inventoryState" bson:"inventoryState,omitempty"`
	DeviceState DeviceState `json:"deviceState" bson:"deviceState,omitempty"`
	Created ChangeLog `json:"created" bson:"created,omitempty"`
	LastModified ChangeLog `json:"lastModified" bson:"lastModified,omitempty"`
}

func (d Device) RequiredFields() []string {
	return []string{"id", "type", "macAddress", "comment", "deleted", "location", "product", "inventoryState", "deviceState"}
}

func NewDevice(devicetype string,
	macaddress string,
	comment string,
	deleted bool,
	
	providerId int, /* Provider */
	providerName string,
	
	aggrZone string, /* Cabinet */
	datacenterId int,
	datacenter string,
	cabinetName string,
	cabinetStartingSpace int,
	
	productCatalogId int, /* Product */
	productCatalogOfferingDesc string,
	productId string,
	productName string,
	
	// inventorystate string, /* Inventory State */
	// inventorycomment string,
	// automationEvent string,
	// accountId string,
	// quoteId string,
	// salesperson string,
	// opportunity string,

	devicestate string, /* Device State */
	deviceComment string,
	userId string) (error, Device) {

	device := Device{}
	if devicetype == "" {
		return ValidationError{"Device type cannot be empty", 400}, device
	}
	device.DeviceType = devicetype
	
	if macaddress == "" {
		return ValidationError{"macaddress cannot be empty", 400}, device
	}
	device.MacAddress = macaddress
	
	if comment == "" {
		return ValidationError{"Comment cannot be empty", 400}, device
	}
	device.Comment = comment

	device.Deleted = deleted

	err, provider := NewProvider(providerId, providerName)
	if err != nil {
		return err, device
	}
	device.Provider = provider

	err, location := NewLocation(aggrZone, datacenterId, datacenter, cabinetName, cabinetStartingSpace)
	if err != nil {
		return err, device
	}
	device.Location = location
	
	err, product := NewProduct(productCatalogId, productCatalogOfferingDesc, productId, productName)
	if err != nil {
		return err, device
	}
	device.Product = product

	device.Id = bson.NewObjectId()
	
	err, deviceState := NewDeviceState(devicestate, device.Id, deviceComment, userId)
	if err != nil {
		return err, device
	}
	
	device.DeviceState = deviceState

	// err, inventoryState := NewInventoryState(inventorystate,
	// 	userId,
	// 	inventorycomment,
	// 	device.Id,
	// 	automationEvent,
	// 	accountId,
	// 	quoteId,
	// 	salesperson,
	// 	opportunity)
	// if err != nil {
	// 	return err, device
	// }
	// device.InventoryState = inventoryState

	device.InventoryState = InventoryState{Device:device.Id}
	
	return nil, device
}

func (d *Device) Save(userId string) error {
	/* TODO: Insert Commend, but which one? */

	if d.Created.TimeStamp.IsZero() {
		d.Created.UserId = userId
		d.Created.TimeStamp = time.Now()
		d.LastModified.UserId = userId
		d.LastModified.TimeStamp = d.Created.TimeStamp
		/* save devicestate */
		if err := d.DeviceState.Save(); err != nil {
			return err
		}
	} else {
		fmt.Print("Old one")
		d.LastModified.UserId = userId
		d.LastModified.TimeStamp = time.Now()
	}

	_, err := backend.GetDB().C("Devices").UpsertId(d.Id, &d)
	return err
}
	
func GetDevices(offset, limit int) []Device {
	var res []Device
	query := backend.GetDB().C("Devices").Find(bson.M{})
	query.Skip(offset).Limit(limit).All(&res)
	return res
}

func GetOneDevice(id string) (error, Device) {
	var res Device
	err := backend.GetDB().C("Devices").FindId(bson.ObjectIdHex("57768f171973d719345e85bb")).One(&res)
	return err, res
}
