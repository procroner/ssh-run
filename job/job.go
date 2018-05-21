package job

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"github.com/remfath/ssh-run/server"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
)

const configPath = "/etc/sshrun/job.json"

const (
	STATUS_DISABLE = iota
	STATUS_ENABLE
	STATUS_RUNNING
	STATUS_STOP
	STATUS_WATING
)

var statusDesc = map[int]string{
	0: "disable",
	1: "enable",
	2: "running",
	3: "stop",
	4: "waiting",
}

type Job struct {
	Name    string `json:"name"`
	Server  string `json:"server"`
	Command string `json:"command"`
	Status  int    `json:"status"`
}

var Jobs map[string]*Job

func init() {
	taskMap, err := parseConfigFile(configPath)
	if err != nil {
		log.Fatalf("parse Job config file error: %v\n", err)
		os.Exit(1)
	}
	Jobs = taskMap
}

func parseConfigFile(configPath string) (map[string]*Job, error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var tasks []*Job
	json.Unmarshal(raw, &tasks)
	serMap := make(map[string]*Job)
	for _, task := range tasks {
		serMap[task.Name] = task
	}
	return serMap, nil
}

func (job *Job) Run() (result string, err error) {
	if job.Status != STATUS_ENABLE {
		return "", errors.New("")
	}

	ser, err := server.Get(job.Server)
	if err != nil {
		log.Fatalf("ser %s not found", job.Server)
		return "", errors.New("server not found")
	}

	if ser.AuthType == "pass" {
		if ser.Pass == "" {
			return server.RunCommandAskPass(ser.User, ser.Host, job.Command)
		}
		return server.RunCommandWithPass(ser.User, ser.Host, ser.Pass, job.Command)
	}
	if ser.AuthType == "key" {
		return server.RunCommandWithKey(ser.User, ser.Host, ser.PrivateKeyPath, job.Command)
	}
	return "", errors.New("")
}

func RunAll() {
	for _, job := range Jobs {
		result, err := job.Run()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
		fmt.Println(result)
	}
}

func List() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Server", "Command", "Status"})
	for _, job := range Jobs {
		ser, err := server.Get(job.Server)
		if err != nil {
			log.Fatalf("ser %s not found", job.Server)
			return
		}
		row := []string{
			job.Name,
			ser.Host,
			job.Command,
			statusDesc[job.Status],
		}
		table.Append(row)
	}
	table.Render()
}
