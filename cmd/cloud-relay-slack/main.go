package main

import (
	"fmt"
	"net/http"
	"os"
	//"encoding/json"
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
    req.ParseForm()
    fmt.Fprint(res, req.Form) //"{\n\"text\": \"African or European?\"\n}")
}
