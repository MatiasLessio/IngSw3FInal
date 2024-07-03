package main

import (
	"backend/app"
)

func main() {
	app.StartDbEngine()
	app.StartRoute()
}
