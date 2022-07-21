package helpers

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func GetProfile(ctx *cli.Context) (Profile, error) {
	profile := ctx.String("profile")
	conf, err := ReadConfig(profile)
	if err != nil {
		return Profile{}, fmt.Errorf(err.Error())
	}
	return conf, nil
}

func SetEnvs(envs map[string]string) {
	for envName, envValues := range envs {
		os.Setenv(envName, envValues)
	}
}

func UnSetEnv(envs map[string]string) {
	for envName, _ := range envs {
		os.Unsetenv(envName)
	}
}
