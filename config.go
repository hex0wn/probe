package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

type configStructure struct {
	LogPath  string	`json:"log_path"`
	IPrefix  string	`json:"ip_prefix"`	
	Rules    []*ruleStructure `json:"rules"`
}

type ruleStructure struct {
	Name         string `json:"name"`
	Listen       string `json:"listen"`
	EnableRegexp bool   `json:"enable_regexp"`
	Targets      []*struct {
		Regexp  string         `json:"regexp"`
		regexp  *regexp.Regexp `json:"-"`
		Address string         `json:"address"`
	} `json:"targets"`
	FirstPacketTimeout uint64 `json:"first_packet_timeout"`
}

var config *configStructure

func init() {
	execFile, err := exec.LookPath(os.Args[0])
    execPath, _ := filepath.Abs(execFile)
    dir := filepath.Dir(execPath)
	buf, err := ioutil.ReadFile(path.Join(dir, "config.json"))
	if err != nil {
		fmt.Printf("failed to load config.json: %s\n", err.Error())
	}

	if err := json.Unmarshal(buf, &config); err != nil {
		fmt.Printf("failed to load config.json: %s\n", err.Error())
	}

	if len(config.Rules) == 0 {
		fmt.Printf("empty rule\n", err.Error())
	}

	for i, v := range config.Rules {
		if err := v.verify(); err != nil {
			fmt.Printf("verity rule failed at pos %d : %s\n", i, err.Error())
		}
	}
}

func (c *ruleStructure) verify() error {
	if c.Name == "" {
		return errors.New("empty name")
	}
	if c.Listen == "" {
		return errors.New("invalid listen address")
	}
	if len(c.Targets) == 0 {
		return errors.New("invalid targets")
	}
	if c.EnableRegexp {
		if c.FirstPacketTimeout == 0 {
			c.FirstPacketTimeout = 5000
		}
	}
	for i, v := range c.Targets {
		if v.Address == "" {
			return errors.New(fmt.Sprintf("invalid address at pos %d\n", i))
		}
		if c.EnableRegexp {
			r, err := regexp.Compile(v.Regexp)
			if err != nil {
				return errors.New(fmt.Sprintf("invalid regexp at pos %d : %s\n", i, err.Error()))
			}
			v.regexp = r
		}
	}
	return nil
}
