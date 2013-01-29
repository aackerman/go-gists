package main

import (
  "fmt"
  "os"
  "time"
  "path"
  "flag"
  "strconv"
  "net/http"
)

var port = flag.Int("port", 8000, "Server port")

func main() {
  flag.Parse()
  root, _ := os.Getwd()
  fmt.Printf("Starting server at http://localhost:%v\n", *port)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    filename := path.Clean(r.URL.String())
    file := root + filename
    fmt.Printf("%s - %s\n", filename, time.Now())
    http.ServeFile(w, r, file)
  })
  http.ListenAndServe(":" + strconv.Itoa(*port), nil)
}
