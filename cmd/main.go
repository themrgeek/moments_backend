package main

import "github.com/themrgeek/moments_backend/pkg/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
