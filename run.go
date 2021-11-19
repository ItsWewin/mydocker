// build linux

package main

import (
	log "github.com/sirupsen/logrus"
	"mydocker/cgroups"
	"mydocker/cgroups/subsystems"
	"mydocker/mycontainer"
	"os"
	"strings"
)

func Run(tty bool, cmdArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := mycontainer.NewParentProcess(tty)
	if parent == nil {
		log.Error("parent is nil")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error("Parent start failed: %v", err)
		return
	}

	cgroupManager := cgroups.NewCgroupManager("my-docker-00100")
	defer cgroupManager.Destroy()
	err := cgroupManager.Set(res)
	if err != nil {
		return
	}

	log.Infof("pid: %d", parent.Process.Pid)

	err = cgroupManager.Apply(parent.Process.Pid)
	if err != nil {
		return
	}

	sendInitCommand(cmdArray, writePipe)
	parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
