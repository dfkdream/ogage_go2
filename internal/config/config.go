/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package config

import (
	"time"
)

type Config struct {
	InputDevices   []string
	BrightnessFile string
	Power          Power
	Command        Command
	JoypadBindings map[string]string
}

type Power struct {
	LongPressDuration time.Duration
}

type Command struct {
	Delay    time.Duration
	Interval time.Duration
}

const DEFAULT_BRIGHTNESS_FILE = "/sys/class/backlight/backlight/brightness"

// TODO: Replace mock config with real one
// TODO: Make repeaters configurable
func Load(path string) (*Config, error) {
	return &Config{
		InputDevices: []string{
			"/dev/input/event0",
			"/dev/input/event1",
			"/dev/input/event2",
		},
		BrightnessFile: DEFAULT_BRIGHTNESS_FILE,
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
	}, nil
}

func (c Config) JoypadBinding(code uint16) string {
	return c.JoypadBindings[eventCodeMap[code]]
}
