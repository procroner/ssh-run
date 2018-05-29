package cmdJob

import (
	"fmt"
	"github.com/procroner/ssh-run/core/coreJob"
)

func runAll() {
	jobs := coreJob.All()
	for _, job := range jobs {
		result, err := job.Run()
		fmt.Printf("\n  %s\n", job.Title)
		fmt.Println("----------------------------------------")
		if err != nil {
			fmt.Printf("  ERROR: %s\n", err)
		} else {
			fmt.Println(result)
		}
		fmt.Printf("\n\n")
	}
}

func Run(jobId int) {
	if jobId == 0 {
		runAll()
	} else {
		job := coreJob.Query(jobId)
		job.Run()
	}
}
