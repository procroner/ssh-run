package coreJob

import (
	"errors"
	"github.com/procroner/ssh-run/core/coreServer"
	"github.com/procroner/ssh-run/core/coreServer/coreConnect"
	"github.com/jinzhu/gorm"
	"github.com/procroner/ssh-run/core/coreDb"
	"time"
	"github.com/procroner/ssh-run/core/coreLog"
	"fmt"
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
	Command  string
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
		return "", errors.New("job is disabled")
	}

	server := coreServer.Query(job.ServerId)

	if server.AuthType == "pass" {
		if server.Pass == "" {
			result, err = coreConnect.RunCommandAskPass(server.User, server.Host, job.Command)
		}
		result, err = coreConnect.RunCommandWithPass(server.User, server.Host, server.Pass, job.Command)
	}
	if server.AuthType == "key" {
		result, err = coreConnect.RunCommandWithKey(server.User, server.Host, server.PrivateKeyPath, job.Command)
	}
	result, err = "", errors.New("")
	endTime := time.Now()
	var runStatus int
	if err != nil {
		runStatus = 1
	}
	log := coreLog.Log{
		JobId:     int(job.ID),
		StartTime: &startTime,
		EndTime:   &endTime,
		Output:    result,
		Error:     fmt.Sprintf("%s", err),
		Status:    runStatus,
	}
	coreLog.Create(log)
	return
}
