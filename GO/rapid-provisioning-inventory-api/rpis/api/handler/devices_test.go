package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"rpis/backends"
	"rpis/config"

	"log"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

func init() {
	config.LoadTestConfiguration()
}

func tearDown() {
	// drop collections used for test cases
	session, err := backends.GetSession()
	if err != nil {
		log.Fatal(err)
	}
	session.DB("").C(config.DeviceCollection).DropCollection()
	session.DB("").C(config.DeviceStateCollection).DropCollection()
}

func TestDeviceRouter(t *testing.T) {
	payload := `[{
		"deviceState": {
			"comment": "dfsdfds",
			"state": "TEST"
		},
		"deleted": false,
		"provider": {
			"id": 9999999,
			"href": "sample link",
			"name": "DELL"
		},
		"location": {
			"cabinetName": "dshkjh",
			"cabinetStartingSpace": null,
			"datacenter": "ORD1",
			"aggrZone": "ORD1:Exnet:Zone7010-TESTDC",
			"datacenterId": 12
		},
		"product": {
			"offerServiceDetails": {
				"href": "offerlink",
				"productId": "3e0bfb6d-dc02-4145-a367-55ebc449a605",
				"productName": "Single Processor Dedicated Server (Haswell)"
			},
			"productCatalogDetails": {
				"offeringDescription": "48GB Dual Processor Hex Core Dedicated Server",
				"href": "productCatalogLink",
				"id": 3025
			}
		},
		"macAddress": "111.111.121",
		"type": "SERVER",
		"comment": "opportunityid=100;account=1550715"
    }]`

	// Create a new device
	router := httprouter.New()
	router.POST("/inventory/devices", CreateDevice)
	req, err := http.NewRequest(
		"POST",
		"/inventory/devices",
		strings.NewReader(payload),
	)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Device not created, HTTP status expected: 201, got: %d", w.Code)
	}

	var s []struct {
		Href string `json:"href"`
		ID   string `json:"id"`
	}

	err = json.Unmarshal(w.Body.Bytes(), &s)
	if err != nil {
		t.Error("Failed to unmarshal output of create", err)
	}

	deviceID := s[0].ID
	uri := "/inventory/devices/" + deviceID

	// Get all devices
	router = httprouter.New()
	router.GET("/inventory/devices", Devices)
	req, err = http.NewRequest("GET", "/inventory/devices", nil)
	if err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		fmt.Print(w.Body.String())
		t.Errorf("Devices Get call, HTTP status expected: 200, got: %d", w.Code)
	}

	// Get created device
	router.GET("/inventory/devices/:id", Device)
	req, err = http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Device Get call, HTTP status expected: 200, got: %d", w.Code)
	}

	// Delete created device
	t.Logf("Deleting %s", uri)
	router.DELETE("/inventory/devices/:id", DeleteDevice)
	req, err = http.NewRequest(
		"DELETE",
		uri,
		strings.NewReader(`{ "comment": "yo", "automationEventId": "12a-131-dr-3232-45"}`),
	)
	if err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("Device found after delete call, HTTP status expected: 204, got: %d", w.Code)
	}

	tearDown()
}

func TestRouterNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.POST("/inventory/devices", CreateDevice)

	// test not allowed
	req, err := http.NewRequest("GET", "/inventory/devices", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if !(w.Code == http.StatusMethodNotAllowed) {
		t.Errorf("NotAllowed handling failed: Code=%d, Header=%v", w.Code, w.Header())
	}
}
