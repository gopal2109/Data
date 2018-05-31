package schema

// Threshold collection
import (
	"rpis/config"
	"rpis/backends"
	"labix.org/v2/mgo/bson"
	log "github.com/Sirupsen/logrus"
)

type ThresholdState struct {
	DatacenterId int `bson:"datacenterId,omitempty" validate:"number,min=1,max=0"`
	DatacenterAbbreviation string `bson:"datacenterAbbreviation,omitempty" validate:"string,min=1,max=0"`
	Warning int `bson:"warning,omitempty" validate:"number,min=1,max=0"`
	Critical int `bson:"critical,omitempty" validate:"number,min=1,max=0"`
}

type DatacenterThresholds map[string]ThresholdState

type Thresholds struct {
	ID             bson.ObjectId `json:"-" bson:"_id,omitempty" validate:"-"`
	Offering Offering `bson:"offering,omitempty"`
	DatacenterThresholds []ThresholdState `bson:"datacenterThresholds"`
}

type Offering struct {
	Href string `bson:"href,omitempty" validate:"string,min=1"`
	OfferingId int `bson:"offeringId,omitempty" validate:"number,min=1,max=0"`
}

// Save the threshold into the db
func (t Thresholds) Save() (string, error) {
	session, err := backends.GetSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	t.ID = bson.NewObjectId()
	collection := session.DB("").C(config.ThresholdsCollection)
	_, err = collection.UpsertId(t.ID, &t)
	if err == nil {
		log.WithError(err).Warn("failed to save latest threshold")
	}
	return t.ID.Hex(), nil
}

// Find Threshold with filters
func (t Thresholds) Find(filters interface{}, offset, limit int) (interface{}, error) {
	session, err := backends.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	collection := session.DB("").C(config.ThresholdsCollection)
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
		var result Thresholds
		err := query.One(&result)
		return result, err
	}
	var results []Thresholds
	err = query.All(&results)
	return results, err
}

// Update an Threshold with all omitempty ignored.
func (t Thresholds) Update(key string, value interface{}) []error {
	u := make(map[string]interface{})
	u[key] = value

	toUpdate := struct {
		Offering Offering `bson:"offering,omitempty"`
		DatacenterThresholds []ThresholdState `bson:"datacenterThresholds"`
	}{
		t.Offering,
		t.DatacenterThresholds,
	}
	errs := ValidateFields("threshold", toUpdate)
	if len(errs) > 0 {
		return errs
	}

	d := map[string]interface{}{"$set": t}
	session, err := backends.GetSession()
	if err != nil {
		return []error{err}
	}
	defer session.Close()

	collection := session.DB("").C(config.ThresholdsCollection)
	if err = collection.Update(u, &d); err != nil {
		return []error{err}
	}
	log.WithFields(log.Fields{
		"key":    key,
		"value":  value,
		"action": "Update",
	}).Info("Threshold is updated")
	return nil
}

func (t Thresholds) Delete(key string, value interface{}) []error {

	return nil
}
