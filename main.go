package main

import (
	"github.com/procroner/ssh-run/job"
	"github.com/procroner/ssh-run/server"
)

func main() {
	server.List()
	job.List()
	job.RunAll()
}
