package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kamaleshnneerasa/GoExamples/handlers1"
)

func main() {
	logObj := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logObj)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	http.ListenAndServe(":9090", serveMux)
}
