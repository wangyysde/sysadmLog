package sysadmlog_test

import (
	"log"
	"net/http"

	"github.com/wangyysde/sysadmlog"
)

func ExampleLogger_Writer_httpServer() {
	logger := sysadmlog.New()
	w := logger.Writer()
	defer w.Close()

	srv := http.Server{
		// create a stdlib log.Logger that writes to
		// sysadmlog.Logger.
		ErrorLog: log.New(w, "", 0),
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func ExampleLogger_Writer_stdlib() {
	logger := sysadmlog.New()
	logger.Formatter = &sysadmlog.JSONFormatter{}

	// Use sysadmlog for standard log output
	// Note that `log` here references stdlib's log
	// Not sysadmlog imported under the name `log`.
	log.SetOutput(logger.Writer())
}
