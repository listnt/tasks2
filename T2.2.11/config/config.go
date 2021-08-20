package config

import (
	"encoding/json"
	"fmt"
	"io"
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
	p, _ := io.ReadAll(jsonFile)
	if err := json.Unmarshal(p, &cfg); err != nil {
		fmt.Println("Wrong config")
		os.Exit(1)
	}
	return cfg
}
