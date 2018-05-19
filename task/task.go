package task

import (
	"io/ioutil"
	"encoding/json"
	"github.com/remfath/ssh-run/connect"
)

type Task struct {
	Name           string `json:"name"`
	User           string `json:"user"`
	Host           string `json:"host"`
	AuthType       string `json:"authType"`
	Pass           string `json:"pass"`
	PrivateKeyPath string `json:"privateKeyPath"`
	Command        string `json:"command"`
}

func Parse(configPath string) ([]*Task, error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var task []*Task
	json.Unmarshal(raw, &task)
	return task, nil
}

func (task *Task) Run() (result string, err error) {
	if task.AuthType == "pass" {
		result, err = connect.RunCommandWithPass(task.User, task.Host, task.Pass, task.Command)
	}
	if task.AuthType == "key" {
		result, err = connect.RunCommandWithKey(task.User, task.Host, task.PrivateKeyPath, task.Command)
	}
	return
}
