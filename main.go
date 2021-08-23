package main

import (
	"uni/controller"
	"uni/repository"
)

func main() {
	repository.Init()

	controller.Run(":2727")
}
