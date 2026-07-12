/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package main

import (
	"log/slog"
	"ogage_go2/internal/config"
	"ogage_go2/internal/evdev"
	"sync/atomic"
	"time"

	"golang.org/x/sys/unix"
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
	config.CmdVolumeUp:       volumeUpRepeater.Stop,
	config.CmdVolumeDown:     volumeDownRepeater.Stop,
}

var hotkeyPressed atomic.Bool

var powerButtonTimer *time.Timer

func handleEvent(event *evdev.InputEvent, dev *evdev.InputDevice) {
	// Handle power button events
	if event.Code == evdev.EVENT_POWER {
		if hotkeyPressed.Load() {
			powerWithHotkey()
			return
		}

		if powerButtonTimer == nil {
			powerButtonTimer = time.AfterFunc(
				config.Get().Power.LongPressDuration,
				powerWithHotkey,
			)
			powerButtonTimer.Stop()
		}

		if event.Value == evdev.VALUE_PRESSED {
			powerButtonTimer.Reset(
				config.Get().Power.LongPressDuration,
			)
		} else {
			powerButtonTimer.Stop()
			power()
		}

		return
	}

	cmd := config.Get().GetCommand(event.Code)

	// Handle hotkey events
	if cmd == config.CmdHotkey {
		if event.Value == evdev.VALUE_PRESSED {
			hotkeyPressed.Store(true)

			if config.Get().Experimental.GrabHotkeyInput {
				err := dev.Grab()
				if err != nil {
					slog.Error(
						"Failed to grab device.",
						"err", err,
						"errno", int(err.(unix.Errno)),
						"dev", dev.File.Name(),
					)
				}
			}
		} else {
			hotkeyPressed.Store(false)

			if config.Get().Experimental.GrabHotkeyInput {
				err := dev.Release()
				if err != nil {
					slog.Error(
						"Failed to release device.",
						"err", err,
						"errno", int(err.(unix.Errno)),
						"dev", dev.File.Name(),
					)
				}

				// TODO: Fix hotkey release issue
				// Currently, if event grab is enabled,
				// hotkey will never be released
				// as release event is grabbed,
				// which makes emulationstation unresponsive.
				// This can be temporarily resolved
				// by removing system_hk from /etc/emulationstation/es_input.cfg.
				// However this issue should be fixed.
				// E.g. injecting key release event after release.
			}
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
