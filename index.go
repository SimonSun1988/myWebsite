package main

import (
    "html/template"
    "net/http"
)

func pageHomeHandler(res http.ResponseWriter, req *http.Request) {
    t, _ := template.ParseFiles("index.html")
    t.Execute(res, nil)
}

func main() {
    http.HandleFunc("/", pageHomeHandler)
    http.ListenAndServe(":9991", nil)
}