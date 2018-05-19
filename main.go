package main

import (
	"github.com/remfath/server-manager/task"
	"fmt"
	"os"
	"os/user"
	"log"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	var configPath = usr.HomeDir + "/ssh_run.json"

	tasks, err := task.Parse(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, task := range tasks {
		result, err := task.Run()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}
