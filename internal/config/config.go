package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	Log        `yaml:"log"`
	Storage    `yaml:"storage"`
	HTTPServer `yaml:"http_server"`
}

type Log struct {
	RequestIDKey string `yaml:"request_id_key" env-default:"request-id"`
}

type Storage struct {
	FsFolderPath string `yaml:"fs_folder_path" env:"FS_FOLDER_PATH" env-default:"/var/tmp/www/safe-concept-server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"127.0.0.1:48080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func findConfigPath() (res string) {
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

func MustLoad() *Config {
	configPath := findConfigPath()
	if configPath == "" {
		panic("config path is not specified")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
