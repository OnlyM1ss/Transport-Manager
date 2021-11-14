package main

import (
	"github.com/OnlyM1ss/transport-manager/v2/app/route"
	"github.com/fatih/color"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	//conf := config.New()
	port := "801" //TODO переделать расположение порта
	color.Cyan("🌏 Server running on localhost:" + port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := route.Routes()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(router)
	http.ListenAndServe(":"+port, handler)
}
