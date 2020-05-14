package logger

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/merlin-foundation/log4j.git/enum"
	"github.com/merlin-foundation/log4j.git/model"
)

var levelLog = enum.LoggerLevel
var logger = model.Logger{}

func respondWithJSON(payload interface{}) []byte {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

//Error func
func Error(Message string) string {
	logger.LevelLog = levelLog.Error
	logger.CreatedAt = formatTime(time.Now())
	logger.Message = Message

	return string(respondWithJSON(logger))
}

//Info func
func Info(Message string) string {
	logger.LevelLog = levelLog.Info
	logger.CreatedAt = formatTime(time.Now())
	logger.Message = Message

	return string(respondWithJSON(logger))
}

//Warn func
func Warn(Message string) string {
	logger.LevelLog = levelLog.Warn
	logger.CreatedAt = formatTime(time.Now())
	logger.Message = Message

	return string(respondWithJSON(logger))
}

//Debug func
func Debug(Message string) string {
	logger.LevelLog = levelLog.Debug
	logger.CreatedAt = formatTime(time.Now())
	logger.Message = Message

	return string(respondWithJSON(logger))
}

//Stdout func
func Stdout(logger string) {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.Println(logger)
}

func formatTime(time time.Time) string {
	return time.Format("02-Jan-2006 03:04:05 PM -0700")
}
