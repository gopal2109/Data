package api

import (
	"fmt"
	"net/http"
	"rpis/config"
	"rpis/api/handlers"
	router "github.com/julienschmidt/httprouter"
)

func APIServer() {
	fmt.Println("Starting server..")
	conf := config.Conf
	
	routes := router.New()

	routes.POST("/inventory/devices", handlers.DeviceHandlerPOST)
	routes.GET("/inventory/devices", handlers.DeviceHandlerGET)
	routes.GET("/inventory/devices/:deviceId", handlers.SingleDeviceHandlerGET)
	routes.PUT("/inventory/devices/:deviceId", handlers.SingleDeviceHandlerPUT)
	routes.DELETE("/inventory/devices/:deviceId", handlers.SingleDeviceHandlerDELETE)
	
	server := http.Server{Addr: conf.Http.HostAddress, Handler: routes}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
