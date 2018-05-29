package coreJob

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/procroner/ssh-run/core/coreDb"
	"time"
	"github.com/procroner/ssh-run/core/coreLog"
	"fmt"
	"github.com/procroner/ssh-run/core/coreServer"
	"github.com/procroner/ssh-run/core/coreServer/coreConnect"
)

const (
	STATUS_DISABLE = iota
	STATUS_ENABLE
	STATUS_RUNNING
	STATUS_STOP
	STATUS_WATING
)

var StatusDesc = map[int]string{
	0: "disable",
	1: "enable",
	2: "running",
	3: "stop",
	4: "waiting",
}

type Job struct {
	gorm.Model
	Title    string
	ServerId int
	Command  string `gorm:"type:longtext"`
	Status   int
}

func Migrate() {
	db := coreDb.Connect()
	defer db.Close()
	db.AutoMigrate(&Job{})
}

func Query(serverId int) Job {
	db := coreDb.Connect()
	defer db.Close()
	var job Job
	db.First(&job, serverId)
	return job
}

func All() []Job {
	db := coreDb.Connect()
	defer db.Close()
	var jobs []Job
	db.Find(&jobs)
	return jobs
}

func (job *Job) Run() (result string, err error) {
	startTime := time.Now()

	if job.Status != STATUS_ENABLE {
		result, err = "", errors.New("job is disabled")
	} else {
		server := coreServer.Query(job.ServerId)

		if server.AuthType == "pass" {
			if server.Pass == "" {
				result, err = coreConnect.RunCommandAskPass(server.User, server.Host, job.Command)
			}
			result, err = coreConnect.RunCommandWithPass(server.User, server.Host, server.Pass, job.Command)
		} else if server.AuthType == "key" {
			result, err = coreConnect.RunCommandWithKey(server.User, server.Host, server.PrivateKeyPath, job.Command)
		} else {
			result, err = "", errors.New("auth type is not allowed")
		}
	}

	endTime := time.Now()
	var runStatus int
	if err != nil {
		runStatus = 1
	}

	var errMsg string
	if err != nil {
		errMsg = fmt.Sprintf("%s", err)
	}

	log := coreLog.Log{
		JobId:     int(job.ID),
		StartTime: &startTime,
		EndTime:   &endTime,
		Output:    result,
		Error:     errMsg,
		Status:    runStatus,
		Command:   job.Command,
	}
	coreLog.Create(log)
	return
}
