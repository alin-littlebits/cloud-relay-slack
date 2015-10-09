package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"io/ioutil"
	//"encoding/json"
)

var hc http.Client
var api_call string = "https://api-http.littlebitscloud.cc/devices/243c200ccecb/output?access_token=97612f7c1ce1b4bd1e2c317bad9f1c4af67e6fb6267931c0"

func main() {
    http.HandleFunc("/", hello)

    hc = http.Client{}

    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    //req.ParseForm()
    //fmt.Fprint(res, req.Form)

    req_req, req_err := http.NewRequest("POST", api_call, nil)
	if req_err != nil {
		log.Fatal(req_err)
	}

	client_req, client_err := hc.Do(req_req)
	if client_err != nil {
		log.Fatal(client_err)
	}

	req_res_body, req_err := ioutil.ReadAll(client_req.Body)
	client_req.Body.Close()

	fmt.Fprintf(res, req_res_body)
}
