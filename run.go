// +build linux

package main

import (
	log "github.com/sirupsen/logrus"
	"mydocker/mycontainer"
)

func Run(tty bool, command string) {
	parent := mycontainer.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	parent.wait()
	os.Exit(-1)
}
