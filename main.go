package main

import (
	"fmt"
	"github.com/mohamedveron/events_detector/api"
	"github.com/mohamedveron/events_detector/service"
	"net/http"
)

func main() {


	//initialize the service layers
	serviceLayer := service.NewService()

	server := api.NewServer(serviceLayer)

	http.HandleFunc("/hello", server.GetAirports)

	fmt.Println("server starting on port 9090...")

	if err := http.ListenAndServe(":9090", nil); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}
