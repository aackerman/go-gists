package main

import (
  "fmt"
  "os"
  "path"
  "net/http"
)

func main() {
  root, _ := os.Getwd()
  fmt.Printf("Starting server at http://localhost:8000\n")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    file := root + path.Clean(r.URL.String())
    http.ServeFile(w, r, file)
  })
  http.ListenAndServe(":8000", nil)
}
