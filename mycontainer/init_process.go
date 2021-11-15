// +build linux

package mycontainer

import (
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func InitProcess(cmd string, args []string) error {
	log.Info("command: %s", cmd)

	defaultMountFlag := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	if err := syscall.Exec(cmd, args, os.Environ()); err != nil {
		log.Errorf("init process comd exec failed  failed")
		return err
	}

	return nil
}
