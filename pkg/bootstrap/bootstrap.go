package bootstrap

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// SetLogLevel ...
func SetLogLevel(level log.Level) {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000",
		DisableColors:   false,
		FullTimestamp:   true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}
