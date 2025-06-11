// Copyright 2025 GEEKROS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

var Get = &Config{}

type Config struct {
	Path      string  `yaml:"-"`
	Workspace string  `yaml:"-"`
	Runtime   string  `yaml:"-"`
	Server    service `yaml:"server"`
}

type service struct {
	Mode         string        `yaml:"mode"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func New() *Config {

	workspace := "/opt/geekros"

	configPath := filepath.Join(workspace, "/release/config.sample.yaml")

	return &Config{
		Path:      configPath,
		Workspace: workspace,
		Runtime:   filepath.Join(workspace, "/runtime"),
	}
}

func (c *Config) LoadConfig() *Config {

	if _, err := os.Stat(c.Path); os.IsNotExist(err) {
		c.Server.Mode = "debug"
		c.Server.Port = 8090
		c.Server.ReadTimeout = 60 * time.Second
		c.Server.WriteTimeout = 60 * time.Second
		c.UpdateConfig()
	}

	file, err := os.ReadFile(c.Path)
	if err != nil {
		return c
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return c
	}

	return c
}

func (c *Config) UpdateConfig() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.Path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
