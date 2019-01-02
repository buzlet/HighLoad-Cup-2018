package main

import "net/http"

func DumbHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("Empty"))
}

func main () {
    http.HandleFunc("/", DumbHandler)
    http.ListenAndServe(":80", nil)
}

