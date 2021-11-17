// build linux

package main

import (
	log "github.com/sirupsen/logrus"
	"mydocker/cgroups"
	"mydocker/cgroups/subsystems"
	"mydocker/mycontainer"
	"os"
)

func Run(tty bool, command string, res *subsystems.ResourceConfig) {
	parent := mycontainer.NewParentProcess(tty, command)

	if parent == nil {
		log.Error("New parent process error")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	log.Infof("[Run] res: %#v", res)

	cgroupManager := cgroups.NewCgroupManager("my-docker-00100")
	defer cgroupManager.Destroy()
	err := cgroupManager.Set(res)
	if err != nil {
		log.Errorf("[Run] cgroup set failed, err: %s", err)
		return
	}

	err = cgroupManager.Apply(parent.Process.Pid)
	if err != nil {
		log.Errorf("[Run] cgroup apply failed, pid: %d", parent.Process.Pid)
		return
	}

	parent.Wait()
	os.Exit(-1)
}
