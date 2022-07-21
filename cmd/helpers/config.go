package helpers

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(profile string) (Profile, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Profile{}, err
	}
	buf, err := ioutil.ReadFile(fmt.Sprintf("%s/.tf_executor/config.yaml", homeDir))
	if err != nil {
		return Profile{}, err
	}

	conf := &Config{}

	err = yaml.Unmarshal(buf, conf)

	if err != nil {
		return Profile{}, fmt.Errorf("in file ~/.tf_executor/config.yaml: %v", err)
	}
	for k := range conf.Profiles {
		if conf.Profiles[k].Name == profile {
			return conf.Profiles[k], nil
		}
	}
	return Profile{}, nil
}
