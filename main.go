package main

import (
	"net/http"
  "fmt"
  "os"
  "io"
)

func handler(rw http.ResponseWriter, req *http.Request) {
  if req.Method == "OPTIONS" {
    rw.Header().Set("Access-Control-Allow-Headers", "accept, authorization, crop")
    rw.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Write([]byte("HELLO"))
  }
  if req.Method == "POST" {
    req.ParseMultipartForm(32 << 20)
    file, handler, err := req.FormFile("file")
    if err != nil {
      fmt.Println(err)
    }
    defer file.Close()
    fmt.Fprintf(rw, "%v", handler.Header)
    f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    io.Copy(f, file)
  }
}

func main() {
	http.HandleFunc("/upload", handler)
	http.ListenAndServe(":8080", nil)
}
