package common

import (
	log "github.com/Sirupsen/logrus"
	"runtime"
)

var Log *log.Entry

func init() {
	Log = log.NewEntry(log.StandardLogger())
}

//init logger with headers
func InitLogger(headers *HttpLib) {
	logger := log.WithFields(log.Fields{
		"x_bunny_request_id": headers.XBunnyRequestId,
		"x_bunny_session_id": headers.XBunnySessionId,
		"x_calling_service":  headers.XCallingService,
		"x_calling_method":   headers.XCallingMethod,
	})

	Log = logger
}

//recover after panic, log traceback
func LogErr() {
	if err := recover(); err != nil {
		trace := make([]byte, 10024)
		count := runtime.Stack(trace, true)
		Log.WithField("error", err).Printf("[error recovered]")
		Log.Printf("Stack trace, lines %d  trace: %s", count, string(trace))
	}
}
