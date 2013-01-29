package main

import "fmt"
import "time"

func even() {
  for i := 0; i < 100; i++ {
    if(i % 2 == 0 && i != 0) {
      fmt.Println("even")
    }
    if(i == 0) {
      fmt.Println("zero")
    }
    time.Sleep(1 * time.Second)
  }
}

func odd() {
  for i := 0; i < 100; i++ {
    if((i + 1) % 2 == 0) {
      fmt.Println("odd")
    }
    time.Sleep(1 * time.Second)
  }
}

func main() {
  go even()
  odd()
}
