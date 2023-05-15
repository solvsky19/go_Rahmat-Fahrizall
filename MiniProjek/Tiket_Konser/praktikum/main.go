package main

import (
	"praktikum/config"
	m "praktikum/middlewares"
	"praktikum/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":3200"))

}
