package errorpk

import (
	log "github.com/sirupsen/logrus"
	"runtime"
)

func LogErr(errMsg string) {
	_, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Error(errMsg)
}
