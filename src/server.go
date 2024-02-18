package main

import     "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello world 1!!!")
}

func main() {
  http.HandleFunc("/test", handler)
  http.ListenAndServe(":80", nil)
}