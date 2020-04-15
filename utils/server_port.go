package utils

import "os"

type server struct{}

var Server server

func (s *server) PortRunning() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}
	return port
}
