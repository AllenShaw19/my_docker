package main

import (
	"my_docker/container"
	"os"

	log "github.com/sirupsen/logrus"
)

// Run function
func Run(tty bool, cmd string) {
	parent := container.NewParentProcess(tty, cmd)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
