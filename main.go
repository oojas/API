package main

import (
	"github.com/joho/godotenv"
	"placementCracker_api/Controllers"
)

func main() {
	godotenv.Load()
	Controllers.Start()

}
