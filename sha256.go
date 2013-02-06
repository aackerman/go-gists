package main

import (
	"log"
	"io"
	"os"
	"fmt"
  "crypto/sha256"
  "flag"
)

func hash256(f io.Reader) {
  hash := sha256.New()
  _, err := io.Copy(hash, f)
  if err != nil {
    log.Fatal(err)
    return
  }
  fmt.Printf("%x\n", hash.Sum(nil))
}

func main() {
	flag.Parse()
	path := flag.Arg(1)
	file, err := os.Open(path)
  if err != nil {
    log.Fatal("File does not exist")
    return
  }
	hash256(file)
}
