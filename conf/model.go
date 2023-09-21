package config

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	MySQL struct {
		DSN string `yaml:"dsn"`
	} `yaml:"MySQL"`
	Admin struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"Admin"`
}
