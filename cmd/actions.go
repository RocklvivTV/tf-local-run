package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rocklvivtv/tf-local-run/cmd/helpers"
	"github.com/urfave/cli/v2"
)

// initTF runs terraform init command
func initTF(ctx *cli.Context) error {
	conf, err := helpers.GetProfile(ctx)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	helpers.SetEnvs(conf.Envs)

	cCmd := exec.Command("terraform", "init")
	cCmd.Stdout = os.Stdout
	cCmd.Stderr = os.Stderr
	cCmd.Dir = ctx.String("path")

	err = cCmd.Run()

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	helpers.UnSetEnv(conf.Envs)
	return nil
}

func planTF(ctx *cli.Context) error {
	conf, err := helpers.GetProfile(ctx)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	helpers.SetEnvs(conf.Envs)

	cCmd := exec.Command("terraform", "plan")
	cCmd.Stdout = os.Stdout
	cCmd.Stderr = os.Stderr
	cCmd.Dir = ctx.String("path")

	err = cCmd.Run()

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	helpers.UnSetEnv(conf.Envs)
	return nil
}

func applyTF(ctx *cli.Context) error {
	conf, err := helpers.GetProfile(ctx)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	helpers.SetEnvs(conf.Envs)

	cCmd := exec.Command("terraform", "apply")
	cCmd.Stdout = os.Stdout
	cCmd.Stderr = os.Stderr
	cCmd.Stdin = os.Stdin
	cCmd.Dir = ctx.String("path")

	err = cCmd.Run()

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	helpers.UnSetEnv(conf.Envs)
	return nil
}

func stateTF(ctx *cli.Context) error {
	return nil
}
