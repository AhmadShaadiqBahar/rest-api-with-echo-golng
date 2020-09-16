package main

import (
	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/db"
	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))

}
