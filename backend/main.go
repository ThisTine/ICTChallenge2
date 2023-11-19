package main

import (
	db "backend/loaders/database"
	"backend/loaders/fiber"
	"backend/loaders/hub"
)

func main() {
	hub.Init()
	db.Init()
	fiber.Init()
}
