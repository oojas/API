package main

import (
	"github.com/joho/godotenv"
	"placementCracker_api/Routers"
)

func main() {
	godotenv.Load()
	Routers.Start()

}
