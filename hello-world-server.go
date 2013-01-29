package main

import (
  "fmt"
  "flag"
  "strconv"
  "net/http"
)

var port = flag.Int("port", 8000, "Server port")

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World!")
}

func main() {
  flag.Parse()
  fmt.Printf("Starting server at %d\n", *port)
  http.HandleFunc("/", handler)
  http.ListenAndServe(":" + strconv.Itoa(*port), nil)
}
