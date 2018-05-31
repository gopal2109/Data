package api

import (
	"fmt"
	"net/http"
	"rpis/config"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// Start called by main to start the service
func Start() {
	address := config.C().Server.Host + ":" + strconv.Itoa(config.C().Server.Port)
	router := getRouter()
	server := http.Server{Addr: address, Handler: router}
	log.WithFields(log.Fields{
		"address": address,
	}).Info(fmt.Sprintf("Serving on %s", address))

	err := server.ListenAndServe()
	if err != nil {
		log.WithError(err).Fatal("Failed to start RPI server")
	}
}
