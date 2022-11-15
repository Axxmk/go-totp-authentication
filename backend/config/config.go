package config

type config struct {
	JWT_SECRET string `yaml:"JWT_SECRET"`
	DB_HOST    string `yaml:"DB_HOST"`
}
