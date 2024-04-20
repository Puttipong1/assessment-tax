package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		output = zerolog.MultiLevelWriter(os.Stderr)

		log = zerolog.New(output).
			With().
			Timestamp().
			Logger()
	})

	return log
}
