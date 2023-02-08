package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	Emojis []*Emoji
	Types  []*Type
}

func LoadConfig() *Config {
	var config Config
	var fns = []func() (string, error){
		os.Getwd,
		os.UserHomeDir,
		os.UserConfigDir,
	}
	for _, fn := range fns {
		if loadConfigByDir(&config, fn) {
			break
		}
	}
	if len(config.Emojis) == 0 {
		config.Emojis = DefaultEmojis[:]
	}
	if len(config.Types) == 0 {
		config.Types = DefaultTypes[:]
	}
	return &config
}

func loadConfigByDir(config *Config, dir func() (string, error)) bool {
	path, err := dir()
	if err != nil {
		return false
	}
	file, err := os.OpenFile(filepath.Join(path, ".gcz.json"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return false
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(file)
	if err != nil && err != io.EOF {
		return false
	}
	var ts Config
	err = json.Unmarshal(buf.Bytes(), &ts)
	if err != nil {
		return false
	}
	*config = ts
	return true
}
