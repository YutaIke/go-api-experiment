package config

import (
	"log"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Configration struct {
	Adderss          string `env:"ADDRESS" envDefault:":1323"`
	Dialect          string `env:"DIALECT" envDefault:"mysql"`
	ConnectionString string `env:"CONNECTION_STRING,required"`
}

func NewConfig(files ...string) (*Configration, error) {
	err := godotenv.Load(files...)
	if err != nil {
		log.Printf(".env file cound be found %q/n", files)
	}

	cfg := Configration{}
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
