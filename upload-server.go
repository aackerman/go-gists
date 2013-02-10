package main

import (
  "io"
  "fmt"
  "net/http"
  "os"
  "path"
  "mime/multipart"
  "encoding/json"
)

type UploadResponse struct {
  Result string `json:"result"`
}

func WriteFileChunk(chunk *multipart.Part, file *os.File) (error) {
  buffer := make([]byte, 4096)
  bufbytes, err := chunk.Read(buffer)
  if err == io.EOF {
    return err
  }
  file.Write(buffer[:bufbytes])
  return err
}

func LargeFileStream(r *multipart.Reader) {
  p, err := r.NextPart()
  if err == io.EOF {
    return
  }

  file, err := os.OpenFile(p.FileName(), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
  if err != nil {
    return
  }
  defer file.Close()

  for {
    err := WriteFileChunk(p, file)
    if err == io.EOF {
      return
    }
  }

  for {
    p, err := r.NextPart()
    if err == io.EOF {
      return
    }

    for {
      err := WriteFileChunk(p, file)
      if err == io.EOF {
        break
      }
    }
  }
  return
}

func upload (w http.ResponseWriter, r *http.Request){
  result := "ok"
  reader, err := r.MultipartReader()
  if err != nil {
    result = "error"
  }

  LargeFileStream(reader)

  // send progress or complete json message
  ur, _ := json.Marshal(UploadResponse{
    Result: result,
  })
  fmt.Fprintf(w, string(ur))
}

func main() {
  root, _ := os.Getwd()

  fmt.Printf("\nStarting server at http://localhost:8000\n\n")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    file := root + path.Clean(r.URL.String())
    http.ServeFile(w, r, file)
  })

  http.HandleFunc("/upload", upload)
  http.ListenAndServe(":8000", nil)
}
