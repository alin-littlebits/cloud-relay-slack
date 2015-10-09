package main

import (
	"fmt"
	"net/http"
	"os"
	//"encoding/json"
)

var hc http.Client
api_call := "https://api-http.littlebitscloud.cc/devices/243c200ccecb/output?access_token=97612f7c1ce1b4bd1e2c317bad9f1c4af67e6fb6267931c0"

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

    client_req, client_err =  http.NewRequest("POST", api_call, nil)

    hc_resp, hc_err := hc.Do(client_req)
    fmt.Fprint(res, hc_resp.Body)
}
