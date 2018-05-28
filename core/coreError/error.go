package coreError

import (
	"fmt"
	"log"
	"os"
)

func HandleError(msg string, err error) {
	if err != nil {
		msg := fmt.Sprintf("[%s]: %v", msg, err)
		fmt.Println(msg)
		log.Fatalln(msg)
		os.Exit(1)
	}
}
