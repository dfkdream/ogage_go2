/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package config

import (
	"ogage_go2/internal/evdev"
	"time"
)

type Config struct {
	InputDevices                 []string
	BrightnessFile               string
	Hotkey                       uint16
	Combinations                 Combinations
	PowerButtonLongPressDuration time.Duration
}

type Combinations struct {
	Delay          time.Duration
	Interval       time.Duration
	BrightnessUp   uint16
	BrightnessDown uint16
	VolumeUp       uint16
	VolumeDown     uint16
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
		Hotkey:         evdev.EVENT_TRIGGER_HAPPY5,
		Combinations: Combinations{
			Delay:          500 * time.Millisecond,
			Interval:       80 * time.Millisecond,
			BrightnessUp:   evdev.EVENT_DPAD_UP,
			BrightnessDown: evdev.EVENT_DPAD_DOWN,
			VolumeUp:       evdev.EVENT_DPAD_RIGHT,
			VolumeDown:     evdev.EVENT_DPAD_LEFT,
		},
		PowerButtonLongPressDuration: 1 * time.Second,
	}, nil
}
