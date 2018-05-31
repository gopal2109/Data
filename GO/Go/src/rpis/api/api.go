package api

import (
	"fmt"
	router "github.com/julienschmidt/httprouter"
	"net/http"
	"rpis/api/handlers"
	"rpis/backend"
	"rpis/config"
)

var log = backend.Log

func APIServer() {

	conf := config.Conf

	routes := router.New()

	routes.POST("/inventory/devices", handlers.DeviceHandlerPOST)
	routes.GET("/inventory/devices", handlers.DeviceHandlerGET)

	routes.GET("/inventory/devices/:deviceId", handlers.SingleDeviceHandlerGET)
	routes.PUT("/inventory/devices/:deviceId", handlers.SingleDeviceHandlerPUT)
	routes.DELETE("/inventory/devices/:deviceId", handlers.SingleDeviceHandlerDELETE)

	routes.GET("/inventory/devices/:deviceId/inventory-states", handlers.InventoryStatesHandlerGET)
	routes.POST("/inventory/devices/:deviceId/inventory-states", handlers.InventoryStatesHandlerPOST)

	routes.GET("/inventory/thresholds", handlers.ThresholdsHandlerGET)
	routes.GET("/inventory/thresholds/:offeringId", handlers.SingleThresholdsHandlerGET)
	routes.PUT("/inventory/thresholds/:offeringId", handlers.ThresholdsHandlerPUT)

	server := http.Server{Addr: conf.Http.HostAddress, Handler: routes}
	log.Info(fmt.Sprintf("Serving on %s", conf.Http.HostAddress))

	err := server.ListenAndServe()
	if err != nil {
		log.Error(err.Error())
	}
}
