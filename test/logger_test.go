package test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/merlin-foundation/log4j/logger"
	"github.com/merlin-foundation/log4j/model"
)

var _logger model.Logger

//LoggerError func testing
func loggerError(t *testing.T) {

	body := logger.Error("Test Message Error")

	if err := json.NewDecoder(strings.NewReader(body)).Decode(&_logger); err != nil {
		log.Fatal(err)
	}

	if _logger.LevelLog != "ERROR" {
		t.Errorf("LoggerLevel :: ERROR different of %s", _logger.LevelLog)
	} else {
		t.Logf("success :: %s", body)
	}
}

//LoggerInfo func testing
func loggerInfo(t *testing.T) {

	body := logger.Info("Test Message Info")

	if err := json.NewDecoder(strings.NewReader(body)).Decode(&_logger); err != nil {
		log.Fatal(err)
	}

	if _logger.LevelLog != "INFO" {
		t.Errorf("LoggerLevel :: INFO different of %s", _logger.LevelLog)
	} else {
		t.Logf("success :: %s", body)
	}
}

//LoggerError func testing
func loggerWarn(t *testing.T) {

	body := logger.Warn("Test Message Warn")

	if err := json.NewDecoder(strings.NewReader(body)).Decode(&_logger); err != nil {
		log.Fatal(err)
	}

	if _logger.LevelLog != "WARN" {
		t.Errorf("LoggerLevel :: WARN different of %s", _logger.LevelLog)
	} else {
		t.Logf("success :: %s", body)
	}
}

//LoggerError func testing
func loggerDebug(t *testing.T) {

	body := logger.Warn("Test Message Debug")

	if err := json.NewDecoder(strings.NewReader(body)).Decode(&_logger); err != nil {
		log.Fatal(err)
	}

	if _logger.LevelLog != "DEBUG" {
		t.Errorf("LoggerLevel :: DEBUG different of %s", _logger.LevelLog)
	} else {
		t.Logf("success :: %s", body)
	}
}
