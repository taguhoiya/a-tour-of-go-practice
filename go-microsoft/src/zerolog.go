package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hey! I'm a log message!")

	log.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "John").
		Send()
}
