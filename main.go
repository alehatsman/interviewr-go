package main

import (
  "github.com/atsman/interviewr-go/db"
  "github.com/atsman/interviewr-go/routes"
)

const (
  Port = "7000"
)

func init()  {
  db.Connect()
}

func main() {
	r := routes.BuildRoutes()

  port := Port
  r.Run(":" + port)
}
