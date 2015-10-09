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
    fmt.Fprint(res, req.PostForm) //"{\n\"text\": \"African or European?\"\n}")
}
