// build linux

package main

import (
	"errors"
	"github.com/urfave/cli"
	"mydocker/cgroups/subsystems"
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
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit (bytes)",
		},
		cli.StringFlag{
			Name:  "cpu",
			Usage: "cpu share limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return errors.New(".Missing container command")
		}

		var cmdArray []string
		for _, arg := range ctx.Args() {
			cmdArray = append(cmdArray, arg)
		}

		tty := ctx.Bool("ti")

		res := &subsystems.ResourceConfig{
			MemoryLimit: ctx.String("m"),
			CpuShare:    ctx.String("cpu"),
			CpuSet:      ctx.String("cpuset"),
		}

		Run(tty, cmdArray, res)
		return nil
	},
}

var initCommand = cli.Command{
	Name:      "init",
	ShortName: "Init container process run user's process in container",
	Action: func(ctx *cli.Context) error {
		err := mycontainer.InitProcess()
		return err
	},
}
