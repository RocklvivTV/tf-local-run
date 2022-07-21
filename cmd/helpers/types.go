package helpers

type Config struct {
	Profiles []Profile `yaml:"profiles"`
}

type Profile struct {
	Name string            `yaml:"name"`
	Envs map[string]string `yaml:"envs"`
}
