package main

import (
	"github.com/remfath/ssh-run/job"
	"github.com/remfath/ssh-run/server"
)

func main() {
	server.List()
	job.List()
	job.RunAll()
}
