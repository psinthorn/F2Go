package configs

import "os"

type server struct{}

var Server server

func (s *server) PortRunning(port string) string {
	envPort := os.Getenv("PORT")
	if envPort == "" {
		envPort = port
	}
	return envPort
}
