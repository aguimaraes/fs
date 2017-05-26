package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	var dir string
	port := flag.String("port", "3000", "port to serve")
	path := flag.String("path", "", "path to serve")

	flag.Parse()

	dir = *path

	if dir == "" {
		dir, _ = os.Getwd()
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))

}
