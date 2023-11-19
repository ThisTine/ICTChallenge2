package main

import (
	"backend/loaders/database"
	"backend/loaders/fiber"
	"backend/loaders/hub"
)

func main() {
	hub.Init()
	database.Init()
	fiber.Init()
}
