package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
)


func main() {
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "{\"text\": \"African or European?\"}")
}
