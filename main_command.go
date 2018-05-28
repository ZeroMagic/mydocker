package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/xianlubird/mydocker/container"
)

var runCommand = cli.Command{
	Name:	"run",
	Usage:	`mydocker run -ti [commond]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:	"ti",
			Usage:	"enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container commond")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		tty := context.Bool("ti")
		Run(tty, cmdArray)
		return nil
	},
}

var initCommand = cli.Command{
	Name:	"init",
	Usage:	"init container",
	Action:	func(context *cli.Context) error {
		log.Info("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}