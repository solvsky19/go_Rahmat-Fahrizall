package main

import (
	"middleware/config"
	"middleware/routes"

	"middleware/middleware"
)

func main() {
	config.InitDB()
	e := routes.New()
	middleware.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))

}
