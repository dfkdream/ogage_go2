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
	"log/slog"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/goccy/go-yaml"
)

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

// TODO: Remove this
const DEFAULT_BRIGHTNESS_FILE = "/sys/class/backlight/backlight/brightness"

var defautConfig = Config{
	InputDevices: []string{
		"/dev/input/event0",
		"/dev/input/event1",
		"/dev/input/event2",
	},
	BrightnessFile: "/sys/class/backlight/backlight/brightness",
	Power: Power{
		LongPressDuration: 1 * time.Second,
	},
	Command: Command{
		Delay:    500 * time.Millisecond,
		Interval: 80 * time.Millisecond,
	},
	JoypadBindings: map[string]string{
		"RIGHT": "VOLUME_UP",
		"LEFT":  "VOLUME_DOWN",
		"UP":    "BRIGHTNESS_UP",
		"DOWN":  "BRIGHTNESS_DOWN",
		"TL":    "VOLUME_DOWN",
		"TR":    "VOLUME_UP",
		"TL2":   "BRIGHTNESS_DOWN",
		"TR2":   "BRIGHTNESS_UP",
		"F5":    "HOTKEY",
	},
}

var globalConfig atomic.Pointer[Config]

// Get returns pointer to the loaded config.
// Do not modify anything on it.
func Get() *Config {
	return globalConfig.Load()
}

func Watch(path string) {

	Load(path)

	// TODO: Make something better
	go func() {
		lastChanged := time.Now()

		for {
			time.Sleep(1 * time.Second)

			info, err := os.Stat(path)
			if err != nil {
				slog.Error(
					"Failed to stat config file. Stopping config watcher.",
					"path", path,
					"err", err,
				)

				break
			}

			modTime := info.ModTime()
			// if lastchanged >= modTime
			if !lastChanged.Before(modTime) {
				continue
			}

			slog.Info(
				"Config file changed. Reloading.",
				"path", path,
				"modTime", modTime,
				"lastChanged", lastChanged,
			)

			lastChanged = modTime

			Load(path)
		}
	}()
}

func Load(path string) {
	f, err := os.Open(path)

	// Create default config file if not exists
	if errors.Is(err, os.ErrNotExist) {
		slog.Info(
			"Config file not found. Creating default one.",
			"path", path,
		)

		err = createDefaultConfig(path)
		if err != nil {
			slog.Error(
				"Failed to create default config file. Falling back to default config.",
				"path", path,
				"err", err,
			)

			globalConfig.Store(&defautConfig)
			return
		}

		f, err = os.Open(path)
	}

	if err != nil {
		slog.Error(
			"Failed to open config file. Falling back to default config.",
			"path", path,
			"err", err,
		)

		globalConfig.Store(&defautConfig)
		return
	}

	defer f.Close()

	var c Config
	err = yaml.NewDecoder(f).Decode(&c)

	if err != nil {
		slog.Error(
			"Failed to decode config. Falling back to default config.",
			"path", path,
			"err", err,
		)

		globalConfig.Store(&defautConfig)
		return
	}

	globalConfig.Store(&c)
}

func (c Config) JoypadBinding(code uint16) string {
	return c.JoypadBindings[eventCodeMap[code]]
}

//go:embed config.yml
var defaultConfigBytes []byte

func createDefaultConfig(path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, defaultConfigBytes, 0666)
	if err != nil {
		return err
	}

	// Use FileMode 0666 for easy editing
	// Chmod is required because of umask.
	err = os.Chmod(path, 0666)
	if err != nil {
		return err
	}

	return nil
}
