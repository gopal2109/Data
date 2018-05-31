package backends

import (
	"rpis/config"

	log "github.com/Sirupsen/logrus"
	"labix.org/v2/mgo"
)

var session *mgo.Session

// GetSession returns a session clone
func GetSession() (*mgo.Session, error) {
	if session == nil {
		mongoURI := config.C().Mongo.ConnectionString
		session, err := mgo.Dial(mongoURI)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"ConnectionString": mongoURI,
			}).Error("Error connecting to Mongo")
		}
		return session, err
	}
	// TODO: need to understand the benchmarks of clone vs copy and it's concurrency benefits.
	return session.Clone(), nil
}
