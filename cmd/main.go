package main

import (
	"assignment-3/router"
)

func main() {
	route := router.Route()

	route.Run(":8080")
}