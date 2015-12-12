package main

import (
	"log"
	"net/http"
	"os"
)

func serve(w http.ResponseWriter, r *http.Request) {
	file := "www" + r.URL.Path
	http.ServeFile(w, r, file)
	log.Print(file)
}

func main() {

	var port = ":8039"
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}
	log.Print("SERVING www on" + port)

	http.HandleFunc("/", serve)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}

}
