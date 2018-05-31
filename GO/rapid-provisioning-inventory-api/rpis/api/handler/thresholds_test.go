package handler

import (
	"fmt"
	"net/http"
	"rpis/backends"
	"strings"
	"net/http/httptest"
	"testing"
	"rpis/api/schema"

	"rpis/config"
	"labix.org/v2/mgo/bson"
	log "github.com/Sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

func init() {
	config.LoadTestConfiguration()
	doc := schema.Thresholds{
		bson.NewObjectId(),
		schema.Offering{"product/catalog/offering/111", 111},
		[]schema.ThresholdState{
			schema.ThresholdState{
			1,"DFW3", 11, 44,
			},
		},

	}
	session, err := backends.GetSession()
	if err != nil {
		log.Fatal(err)
	}

	collection := session.DB("").C(config.ThresholdsCollection)
	errs := collection.Insert(&doc)
	if errs != nil {
		log.Fatal(errs)
	}
}

func CleanupDb() {
	// drop collections used for test cases
	session, err := backends.GetSession()
	if err != nil {
		log.Fatal(err)
	}
	session.DB("").C(config.ThresholdsCollection).DropCollection()
	defer session.Close()
}

func TestThresholdRouter(t *testing.T) {

	offeringId := 111

	uri := fmt.Sprintf("/inventory/thresholds/%d", offeringId)

	// Get all thresholds
	router := httprouter.New()
	router.GET("/inventory/thresholds", ThresholdsHandlerGET)
	req, err := http.NewRequest("GET", "/inventory/thresholds", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Thresholds Get call, HTTP status expected: 200, got: %d", w.Code)
	}

	// Get one threshold
	router.GET("/inventory/thresholds/:offeringId", SingleThresholdsHandlerGET)
	req, err = http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Threshold Get call, HTTP status expected: 200, got: %d", w.Code)
	}

	// upadte threshold
	router.PUT("/inventory/thresholds/:offeringId", UpdateThreshold)
	dict := `{
		"offering": {
			"href": "product/catalog/offering/111",
			"offeringId": 111
		},
		"datacenterThresholds": [{
			"datacenterId": 1,
			"datacenterAbbreviation": "DFW3",
			"warning": 11,
			"critical": 44
		}, {
			"datacenterId": 123,
			"datacenterAbbreviation": "IAD3",
			"warning": 5005,
			"critical": 12
		}, {
			"datacenterId": 777,
			"datacenterAbbreviation": "LON5",
			"warning": 888,
			"critical": 5000
		}]
	}`
	req, err = http.NewRequest("PUT", uri, strings.NewReader(dict),)
	if err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusAccepted {
		t.Errorf("Threshold update call, HTTP status expected: 202, got: %d", w.Code)
	}

	CleanupDb()
}

func TestThresholdRouterNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.GET("/inventory/thresholds", CreateDevice)

	// test not allowed
	req, err := http.NewRequest("POST", "/inventory/thresholds", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if !(w.Code == http.StatusMethodNotAllowed) {
		t.Errorf("NotAllowed handling failed: Code=%d, Header=%v", w.Code, w.Header())
	}
}
