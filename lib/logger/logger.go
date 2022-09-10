package logger

import "github.com/rs/zerolog"

func InitialLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
