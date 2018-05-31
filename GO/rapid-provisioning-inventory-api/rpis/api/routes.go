package api

import (
	"rpis/api/handler"

	"github.com/julienschmidt/httprouter"
)

func getRouter() *httprouter.Router {
	router := httprouter.New()

	// Device handlers
	router.GET("/inventory/devices", handler.RequestLogger(handler.Devices))
	router.POST("/inventory/devices", handler.RequestLogger(handler.CreateDevice))
	router.GET("/inventory/devices/:id", handler.RequestLogger(handler.Device))
	router.DELETE("/inventory/devices/:id", handler.RequestLogger(handler.DeleteDevice))
	router.PUT("/inventory/devices/:id", handler.RequestLogger(handler.UpdateDevice))
	// Thresholds handlers
	router.GET("/inventory/thresholds", handler.RequestLogger(handler.ThresholdsHandlerGET))
	router.GET("/inventory/thresholds/:offeringId", handler.RequestLogger(handler.SingleThresholdsHandlerGET))
	router.PUT("/inventory/thresholds/:offeringId", handler.RequestLogger(handler.UpdateThreshold))

	return router
}
