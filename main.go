package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kamaleshnneerasa/GoExamples/handlers"
)

func main() {
	logObj := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logObj)

	serveMux = http.NewServerMux()
	serveMux.Handle("/", helloHandler)

	http.ListenAndServe(":9090", serveMux)
}
