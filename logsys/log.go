package logsys

import (
	"github.com/rs/zerolog"
)

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
