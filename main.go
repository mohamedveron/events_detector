package main

import (
	"fmt"
	"github.com/mohamedveron/events_detector/api"
	"github.com/mohamedveron/events_detector/service"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {


	//initialize the service layers
	serviceLayer := service.NewService()

	server := api.NewServer(serviceLayer)

	http.HandleFunc("/", entry)
	http.HandleFunc("/event", server.AddEvent)

	fmt.Println("server starting on port 9090...")

	if err := http.ListenAndServe(":9090", nil); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}

func entry(reswt http.ResponseWriter, req *http.Request) {
	tmpl.ExecuteTemplate(reswt, "index.html", nil)

}
