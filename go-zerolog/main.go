package main

import (
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
)

// Filter in Grafana: {container="go-zerolog"}|json|level="warn"|message="hello world"
func main() {
	var randomLogLevel int
	for {
		randomLogLevel = rand.Intn(3)
		switch randomLogLevel {
		case 0:
			log.Print("hello world")
		case 1:
			log.Warn().Msg("hello world")
		case 2:
			log.Error().Msg("hello world")
		}

		time.Sleep(time.Second * 5)
	}
}
