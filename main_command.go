// +build linux

package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"mydocker/mycontainer"
)

var runCommand = cli.Command{
	Name:      "run",
	ShortName: "create a container with namespace and cgroup limit,\n mydocker run -it [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return errors.New(".Missing container command")
		}

		cmd := ctx.Args().Get(0)
		tty := ctx.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:      "init",
	ShortName: "Init container process run user's process in container",
	Action: func(ctx *cli.Context) error {
		log.Infof("init command")
		cmd := ctx.Args().Get(0)
		log.Infof("command: %s", cmd)
		err := mycontainer.InitProcess(cmd, nil)
		return err
	},
}
