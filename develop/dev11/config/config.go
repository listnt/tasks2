package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Config ...
type Config struct {
	ConnectionString string `json:"db_conn"`
}

// ParseConfig ...
func ParseConfig(path string) Config {
	var cfg Config
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	configBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(configBytes, &cfg); err != nil {
		fmt.Println("Wrong config")
		os.Exit(1)
	}
	return cfg
}
