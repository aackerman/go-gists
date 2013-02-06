package main

import (
  "io"
  "fmt"
  "net/http"
  "log"
  "os"
  "path"
  "mime/multipart"
)

func writeMultipart(reader *multipart.Reader, path string) (filepath string, err error) {
  // open a file
  outfile, err := os.Create(path)
  if err != nil {
    return "", err
  }
  defer outfile.Close()

  // loop over the parts of the request
  for {
    // get the next chunk of the file
    chunk, err := reader.NextPart()
    if err == io.EOF {
      break
    }

    // read the chunk in smaller 4KB chunks
    for {
      buffer := make([]byte, 4096)
      bufbytes, err := chunk.Read(buffer)
      if err == io.EOF {
          break
      }
      // write the number of bufbytes to the file
      outfile.Write(buffer[:bufbytes])
    }
  }
  return path, nil
}

func upload (w http.ResponseWriter, r *http.Request){
  // create a reader
  reader, err := r.MultipartReader()
  if err != nil {
    log.Fatal(err);
  }

  writeMultipart(reader, "outfile")
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
