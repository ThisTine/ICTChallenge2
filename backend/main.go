package main

import (
	"backend/loaders/db"
	"backend/loaders/fiber"
	"backend/loaders/hub"
)

func main() {
	db.Init()
	hub.Init()
	fiber.Init()
}
