package coreLog

import (
	"github.com/procroner/ssh-run/core/coreDb"
	"github.com/jinzhu/gorm"
	"time"
)

type Log struct {
	gorm.Model
	JobId     int
	StartTime *time.Time
	EndTime   *time.Time
	Output    string
	Error     string
	Input     string
	Status    int
}

func Migrate() {
	db := coreDb.Connect()
	defer db.Close()
	db.AutoMigrate(&Log{})
}

func Query(logId int) Log {
	db := coreDb.Connect()
	defer db.Close()
	var log Log
	db.First(&log, logId)
	return log
}

func All() []Log {
	db := coreDb.Connect()
	defer db.Close()
	var logs []Log
	db.Find(&logs)
	return logs
}

func Create(log Log) {
	db := coreDb.Connect()
	defer db.Close()
	db.Create(&log)
}