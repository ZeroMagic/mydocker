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
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:	"init",
	Usage:	"init container",
	Action:	func（context *cli.Context) error {
		log.Info("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}