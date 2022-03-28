package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

/*Build info*/
var (
	ReleaseVersion = "Development"
	BuildTime      string
)

func main() {

	port, ver := processFlags()

	if ver {
		fmt.Printf("Version:\t%s\n", ReleaseVersion)
		fmt.Printf("BuildTime:\t%s\n", BuildTime)
		return
	}

	fmt.Println("Starting web server")
	http.HandleFunc("/", webpage)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func webpage(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		host = "UNKNOWN"
	}
	fmt.Fprintf(w, "<h1>This is my basic web page</h1><br>I hope you like it<br>I'm running on %s", host)
}

func processFlags() (uint, bool) {
	var port uint
	var ver bool
	flag.UintVar(&port, "p", 8000, "the port that the webserver will listen on, defaults to 8000 if not set")
	flag.BoolVar(&ver, "v", false, "prints the version and quit")

	flag.Parse()
	return port, ver
}
