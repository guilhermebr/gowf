package gowf

import (
	"os"
	"path/filepath"
)

type Config struct {
	StaticDir string
	HttpAddr  string
	HttpPort  int
	AppPath   string
	workPath  string
}

func InitConfig() *Config {
	workPath, _ := os.Getwd()
	workPath, _ = filepath.Abs(workPath)
	AppPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return &Config{
		StaticDir: "/static",
		workPath:  workPath,
		AppPath:   AppPath,
		HttpAddr:  "127.0.0.1",
		HttpPort:  5000,
	}
}
