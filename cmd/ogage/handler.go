/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package main

import (
	"ogage_go2/internal/config"
	"ogage_go2/internal/evdev"
	"sync/atomic"
	"time"
)

var pressCommands = map[string]func(){
	config.CmdBrightnessUp: func() {
		brightnessUpRepeater.Start(
			config.Get().Command.Delay,
			config.Get().Command.Interval,
		)
	},
	config.CmdBrightnessDown: func() {
		brightnessDownRepeater.Start(
			config.Get().Command.Delay,
			config.Get().Command.Interval,
		)
	},
	config.CmdBrightnessMax: brightnessMax,
	config.CmdBrightnessMin: brightnessMin,
	config.CmdVolumeUp: func() {
		volumeUpRepeater.Start(
			config.Get().Command.Delay,
			config.Get().Command.Interval,
		)
	},
	config.CmdVolumeDown: func() {
		volumeDownRepeater.Start(
			config.Get().Command.Delay,
			config.Get().Command.Interval,
		)
	},
}

var releaseCommands = map[string]func(){
	config.CmdBrightnessUp:   brightnessUpRepeater.Stop,
	config.CmdBrightnessDown: brightnessDownRepeater.Stop,
	config.CmdBrightnessMax:  brightnessMax,
	config.CmdBrightnessMin:  brightnessMin,
	config.CmdVolumeUp:       volumeUpRepeater.Stop,
	config.CmdVolumeDown:     volumeDownRepeater.Stop,
}

var hotkeyPressed atomic.Bool

var powerButtonTimer *time.Timer

func handleEvent(event *evdev.InputEvent) {
	// Handle power button events
	if event.Code == evdev.EVENT_POWER {
		if hotkeyPressed.Load() {
			powerWithHotkey()
			return
		}

		if event.Value == evdev.VALUE_PRESSED {
			// Initialize and start timer
			if powerButtonTimer == nil {
				powerButtonTimer = time.AfterFunc(
					config.Get().Power.LongPressDuration,
					powerWithHotkey,
				)
			} else {
				powerButtonTimer.Reset(
					config.Get().Power.LongPressDuration,
				)
			}
		} else {
			if powerButtonTimer != nil {
				powerButtonTimer.Stop()
			}
			power()
		}

		return
	}

	cmd := config.Get().GetCommand(event.Code)

	// Handle hotkey events
	if cmd == config.CmdHotkey {
		if event.Value == evdev.VALUE_PRESSED {
			hotkeyPressed.Store(true)
		} else {
			hotkeyPressed.Store(false)
		}
		return
	}

	// Handle other events
	if hotkeyPressed.Load() {
		if event.Value == evdev.VALUE_PRESSED {
			if f, ok := pressCommands[cmd]; ok {
				f()
				return
			}
		} else {
			if f, ok := releaseCommands[cmd]; ok {
				f()
				return
			}
		}
	}

	// Process unhandled events here
}
