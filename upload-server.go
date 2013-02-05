package main

import (
  "io"
  "fmt"
  "net/http"
  "log"
  "os"
  "path"
)

func upload (w http.ResponseWriter, r *http.Request){
  fmt.Printf("request %v \n", r)
  // create a reader
  reader, err := r.MultipartReader()
  if err != nil {
    log.Fatal(err);
  }

  // get the length of the request
  length := r.ContentLength

  // loop over the parts of the request
  for {

    // get the next chunk of the file
    chunk, err := reader.NextPart()
    if err == io.EOF {
      break
    }

    // open a file to write to
    outfile, err := os.OpenFile("outfile", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
      log.Fatal(err)
      return
    }
    defer outfile.Close()

    // read the chunk in parts?
    for {
      buffer := make([]byte, 4096)
      bytesread, err := chunk.Read(buffer)
      if err == io.EOF {
          break
      }
      // write the number of bytesread to the file
      outfile.Write(buffer[:bytesread])
    }
  }
}

func main() {
  root, _ := os.Getwd()

  fmt.Printf("Starting server at http://localhost:8000\n")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    file := root + path.Clean(r.URL.String())
    http.ServeFile(w, r, file)
  })

  http.HandleFunc("/upload", upload)
  http.ListenAndServe(":8000", nil)
}
