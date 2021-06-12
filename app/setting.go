package app

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	// Coin Market Cap
	Url     string `yaml:"url" env:"URL"`
	ApiKey  string `yaml:"api_key" env:"API_KEY"`
	ApiHost string `yaml:"api_host" env:"API_HOST"`
}

func Init() {
	log.Println("Setup Environment")
	// Setup Config
	err := cleanenv.ReadConfig(ConfigFile, &Setting)
	if err != nil {
		log.Fatalf("- could not read the config file. %v", err.Error())
	}
	// Setup Service
	Run = &Service{api: &API{}}
	// Setup Agent
	Agent = &Scheduler{
		Service: Run,
	}
}
