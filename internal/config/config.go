/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package config

import (
	_ "embed"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/goccy/go-yaml"
)

//go:embed config.yml
var defaultConfig []byte

type Config struct {
	InputDevices   []string          `yaml:"InputDevices"`
	BrightnessFile string            `yaml:"BrightnessFile"`
	Power          Power             `yaml:"Power"`
	Command        Command           `yaml:"Command"`
	JoypadBindings map[string]string `yaml:"JoypadBindings"`
}

type Power struct {
	LongPressDuration time.Duration `yaml:"LongPressDuration"`
}

type Command struct {
	Delay    time.Duration `yaml:"Delay"`
	Interval time.Duration `yaml:"Interval"`
}

const DEFAULT_BRIGHTNESS_FILE = "/sys/class/backlight/backlight/brightness"

func Load(path string) (*Config, error) {
	f, err := os.Open(path)

	// Create default config file if not exists
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(path, defaultConfig, 0666)
		if err != nil {
			return nil, err
		}

		// Use FileMode 0666 for easy editing
		// Chmod is required because of umask.
		err = os.Chmod(path, 0666)
		if err != nil {
			return nil, err
		}

		f, err = os.Open(path)
	}

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var c Config
	err = yaml.NewDecoder(f).Decode(&c)

	return &c, err
}

func (c Config) JoypadBinding(code uint16) string {
	return c.JoypadBindings[eventCodeMap[code]]
}
