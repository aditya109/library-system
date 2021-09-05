package logger

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Event struct {
	id      int
	message string
}

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

var (
	issueMessage = Event{
		1, "Something went wrong: %s",
	}
	serverEventMessage = Event{
		2, "Server event: %s",
	}
	invalidArgMessage = Event{
		3, "Invalid arg: %s",
	}
	invalidArgValueMessage = Event{
		4, "Invalid value for argument: %s: %v",
	}
	missingArgMessage = Event{
		5, "Missing arg: %s",
	}
	databaseEventMessage = Event{
		6, "Database event: %s",
	}
)

func (l *StandardLogger) Issue(argumentName string) {
	l.Errorf(issueMessage.message, argumentName)
}

func (l *StandardLogger) ServerEvent(argumentName string) {
	l.Errorf(serverEventMessage.message, argumentName)
}

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}

func (l *StandardLogger) DatabaseEvent(argumentName string) {
	l.Errorf(databaseEventMessage.message, argumentName)
}

func RaiseAlert(w http.ResponseWriter, message string, status int) {
	log := NewLogger()
	w.WriteHeader(http.StatusInternalServerError)
	log.Issue(message)
	fmt.Fprintf(w, "%s", message)
}
