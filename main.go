package main

import (
  "github.com/atsman/interviewr-go/db"
  "github.com/atsman/interviewr-go/handlers"
)

const (
  Port = "3000"
)

func init()  {
  db.Connect()
}

func main() {
	r := handlers.NewEngine()

  port := Port
  r.Run(":" + port)
}
