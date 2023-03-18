package util

import (
	log "github.com/sourcegraph-ce/logrus"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var myLogger *logger

var once sync.Once

// start loggeando
func GetLoggerInstance() *logger {
	once.Do(func() {
		myLogger = createLogger("vthunder.log")
	})
	return myLogger
}

func createLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "My app Name ", log.Lshortfile),
	}
}
