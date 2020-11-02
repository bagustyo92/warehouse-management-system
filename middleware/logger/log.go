package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func openLogFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func MakeLogEntry(c echo.Context) *log.Entry {
	// Custom Formater
	customFormater := new(log.TextFormatter)
	customFormater.DisableTimestamp = false
	customFormater.TimestampFormat = "[02-01-2006][15:04:05]"
	customFormater.ForceColors = true

	// Write log into a file, and change filename everyday
	filename := fmt.Sprintf("logs/%s_log.txt", time.Now().Format("02-01-2006"))
	logFile, err := openLogFile(filename)
	if err != nil {
		return log.WithFields(log.Fields{
			"WARNING": err,
		})
	}

	// Write multiple writer into printout and to a file
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set all the misc
	log.SetFormatter(customFormater)
	log.SetOutput(multiWriter)

	if c == nil {
		return log.WithFields(log.Fields{
			"Time": time.Now().Format("[02-01-2006][15:04:05]"),
		})
	}

	return log.WithFields(log.Fields{
		"Method":   c.Request().Method,
		"URL":      c.Request().URL.String(),
		"IP":       c.RealIP(),
		"Lenght":   c.Request().ContentLength,
		"Protocol": c.Request().Proto,
		"Type":     c.Request().Header.Get("Content-Type"),
		// "Time":   time.Now().Format("[02-01-2006][15:04:05]"),
	})
}

func Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Info("Incoming Request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	MakeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
