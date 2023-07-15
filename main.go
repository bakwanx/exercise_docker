package main

import (
	"exercise_docker/config"
	"exercise_docker/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}
