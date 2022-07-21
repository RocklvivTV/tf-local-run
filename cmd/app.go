package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name:  "tf_local",
		Usage: "Helps to run terraform locally with different credentials.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "profile",
				Usage:    "Name of profile with credentials to use",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "path",
				Usage:    "Terraform directory path to use",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "init",
				Usage:   "Runs terraform init in select directory",
				Aliases: []string{"i"},
				Action:  initTF,
			},
			&cli.Command{
				Name:    "plan",
				Usage:   "Runs terraform plan in selected directory",
				Aliases: []string{"p"},
				Action:  planTF,
			},
			&cli.Command{
				Name:    "apply",
				Usage:   "Runs terraform apply in selected directory",
				Aliases: []string{"a"},
				Action:  applyTF,
			},
			&cli.Command{
				Name:    "state",
				Usage:   "Runs terraform state in seleted directory",
				Aliases: []string{"s"},
				Action:  stateTF,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
