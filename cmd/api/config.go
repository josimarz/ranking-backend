package main

import (
	"flag"
	"os"
	"strconv"
)

type config struct {
	port int
}

func loadConfig() (cfg config) {
	flag.IntVar(&cfg.port, "port", getPort(8080), "API server port")
	flag.Parse()
	return
}

func getPort(def int) int {
	if s, exists := os.LookupEnv("_LAMBDA_SERVER_PORT"); exists {
		port, err := strconv.Atoi(s)
		if err == nil {
			return port
		}
	}
	return def
}
