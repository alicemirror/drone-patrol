package webserver

import (
	"arduino/bhy/static"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/browser"
)

func errCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func startServer() {
	fmt.Printf("Starting server at port 8000\n")

	fileServer := http.FileServer(http.FS(static.Content))
	http.Handle("/", fileServer)
	http.ListenAndServe(":8000", nil)
}

func Execute() {
	go startServer()

	e := browser.OpenURL("http://localhost:8000/sensor.html")

	errCheck(e)

	select {}
}
