package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	LogLevel        string `json:"log_level"`
	DBConnectString string `json:"db_conn"`
	ConfigPath      string
}

// Инициализация конфигурации
func NewConfig() (*Config, error) {
	config := &Config{}

	flagLogLevel := flag.String("l", "", "log level")
	flagConfigPath := flag.String("c", "", "config path")
	flagDBConnectString := flag.String("d", "", "db connection string")
	flag.Parse()

	config.ConfigPath = *flagConfigPath

	if configPathEnv := os.Getenv("CONFIG"); configPathEnv != "" {
		config.ConfigPath = configPathEnv
	}

	configFromFile, err := config.ReadConFile(config.ConfigPath)
	if err != nil {
		return &Config{}, err
	}

	config.LogLevel = priorityString(os.Getenv("LOG_LEVEL"), *flagLogLevel, configFromFile.LogLevel, "info")
	config.DBConnectString = priorityString(os.Getenv("DATABASE_DSN"), *flagDBConnectString, configFromFile.DBConnectString)

	return config, nil
}

// Вывод конфигурации в JSON
func (c Config) Json() json.RawMessage {
	b, err := json.Marshal(&c)
	if err != nil {
		return nil
	}
	return b
}

// Выбор первой не пустой строки по порядку приоритета
func priorityString(vars ...string) string {
	for _, v := range vars {
		if v != "" {
			return v
		}
	}
	return ""
}

// Чтение файла конфигурации
func (c *Config) ReadConFile(path string) (Config, error) {
	if path == "" {
		return Config{}, nil
	}

	fileConfig, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, fmt.Errorf("can not read config from - %s", path)
		}
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(fileConfig, &config)
	return config, err
}
