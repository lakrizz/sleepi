package main

import (
	"./lib/http"
)

func main() {
	// Initialize Server
	s, err := http.NewServer("8080")
	if err != nil {
		panic(err)
	}

	// Load Modules and add them
	s.AddModule(testmodule.TestModule)

	// Start Server
	s.Start()
}
