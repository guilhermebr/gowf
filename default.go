package gowf

import (
	"os"
	"path/filepath"
)

var (
	AppPath   string
	workPath  string
	StaticDir map[string]string
	HttpAddr  string
	HttpPort  int
)

func init() {
	workPath, _ = os.Getwd()
	workPath, _ = filepath.Abs(workPath)
	AppPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	StaticDir = make(map[string]string)
	StaticDir["/static"] = "static"
	HttpAddr = ""
	HttpPort = 5000
}
