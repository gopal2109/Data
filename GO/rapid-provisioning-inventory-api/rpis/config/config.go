package config

import (
	"os"
	"sync"
	"time"

	"encoding/json"

	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

const (
	// DeviceCollection is name of the mongo collection where device are stored.
	DeviceCollection = "devices"
	// DeviceStateCollection is name of the mongo collection where deviceStates are stored.
	DeviceStateCollection = "deviceStates"
	// ThresholdsCollection is name of the mongo collection where thresholds are stored.
	ThresholdsCollection = "thresholds"
)

var (
	conf               *Configuration
	configLock         = new(sync.RWMutex)
	configLastReadTime time.Time
	configFilePath     string
)

// Configuration hold service wide information
type Configuration struct {
	Server struct {
		Host  string
		Port  int
		Debug bool
	}
	Mongo struct {
		ConnectionString string
	}
	Logging struct {
		File  string
		Level string
	}
	Application struct {
		TimeZone   string
		DateFormat string
		MaxResults int
	}
}

func readHelper(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.WithError(err).Fatal("Error opening config file")
		os.Exit(1)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.WithError(err).Fatal("Failed while reading config file")
	}
	return buf
}

// Init is the first function that should be called
func Init(configFile string) {
	configFilePath = configFile
	buf := readHelper(configFile)
	c, err := loadConfiguration(buf)
	if err != nil {
		log.WithError(err).Fatal("loadConfiguration failed")
	}
	conf = c
}

// C is the access point for configuration
func C() Configuration {
	configLock.Lock()
	defer configLock.Unlock()
	return *conf
}

func loadConfiguration(data []byte) (*Configuration, error) {
	c := new(Configuration)
	err := json.Unmarshal(data, c)
	if err != nil {
		return c, err
	}
	return c, nil
}

// LoadTestConfiguration loads configuration for test cases
func LoadTestConfiguration() {
	rawConf := []byte(`{
		"Server": {
			"Port": 8082,
			"Debug": true
		},
		"Mongo": {
			"ConnectionString": "mongodb://localhost/rpis_test",
			"CollectionNames": {
				"devices": "devices_test",
				"deviceState": "deviceStates_test",
				"inventoryState": "inventoryStates_test",
				"thresholds": "thresholds_test"
			}
		},
		"Application": {
			"MaxResult": 10
		}
	}`)

	c, err := loadConfiguration(rawConf)
	if err != nil {
		log.WithError(err).Fatal("loadConfiguration failed")
		return
	}

	conf = c
}

// Watcher looks for changes in config file
func Watcher() {
	log.Infof("watching '%s' for change", configFilePath)
	ticker := time.NewTicker(time.Second * 20)
	go func() {
		for {
			<-ticker.C
			if configFilePath != "" {
				f, err := os.Open(configFilePath)
				if err != nil {
					log.WithError(err).Fatal("error opening config file")
				}
				defer f.Close()

				finfo, err := f.Stat()
				if err != nil {
					log.WithError(err).Warn("failed to `stat` file")
				}

				if finfo.ModTime().After(configLastReadTime) {
					buf := readHelper(configFilePath)
					c, err := loadConfiguration(buf)
					if err != nil {
						log.WithError(err).Error("failed to load Configuration")
					} else {
						configLock.Lock()
						conf = c
						configLock.Unlock()
						configLastReadTime = time.Now()
						log.WithField("config_reloaded", "true").Info("Config reloaded")
					}
				}
			} else {
				log.WithFields(log.Fields{
					"configFilePath": configFilePath,
				}).Error("configuration file not specified")
			}
		}
	}()
}
