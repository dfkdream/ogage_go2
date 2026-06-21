/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package main

import (
	"ogage_go2/internal/evdev"
	"ogage_go2/internal/eventprocessor"
	"sync/atomic"
)

var hotkeyPressed atomic.Bool

func powerButtonProcessor(event *evdev.InputEvent) int {
	if event.Code == evdev.EVENT_POWER && event.Value == evdev.VALUE_PRESSED {
		if hotkeyPressed.Load() {
			powerWithHotkey()
		} else {
			power()
		}

		return eventprocessor.HANDLER_ABORT
	}

	return eventprocessor.HANDLER_OK
}

func headphoneProcessor(event *evdev.InputEvent) int {
	if event.Code == evdev.EVENT_HEADPHONE_INSERT {
		if event.Value == evdev.VALUE_INSERTED {
			audioHeadphone()
		} else {
			audioSpeaker()
		}

		return eventprocessor.HANDLER_ABORT
	}

	return eventprocessor.HANDLER_OK
}

func hotkeyProcessor(event *evdev.InputEvent) int {
	if event.Code == conf.Hotkey {
		if event.Value == evdev.VALUE_PRESSED {
			hotkeyPressed.Store(true)
		} else {
			hotkeyPressed.Store(false)

			brightnessUpRepeater.Stop()
			brightnessDownRepeater.Stop()
			volumeUpRepeater.Stop()
			volumeDownRepeater.Stop()
		}

		return eventprocessor.HANDLER_ABORT
	}

	return eventprocessor.HANDLER_OK
}

func combinationPressProcessor(event *evdev.InputEvent) int {
	if !(hotkeyPressed.Load() && event.Value == evdev.VALUE_PRESSED) {
		return eventprocessor.HANDLER_OK
	}

	switch event.Code {
	case conf.Combinations.BrightnessUp:
		brightnessUpRepeater.Start()

	case conf.Combinations.BrightnessDown:
		brightnessDownRepeater.Start()

	case conf.Combinations.VolumeUp:
		volumeUpRepeater.Start()

	case conf.Combinations.VolumeDown:
		volumeDownRepeater.Start()

	default:
		return eventprocessor.HANDLER_OK
	}

	return eventprocessor.HANDLER_ABORT
}

func combinationReleaseProcessor(event *evdev.InputEvent) int {
	if !(hotkeyPressed.Load() && event.Value == evdev.VALUE_RELEASED) {
		return eventprocessor.HANDLER_OK
	}

	switch event.Code {
	case conf.Combinations.BrightnessUp:
		brightnessUpRepeater.Stop()

	case conf.Combinations.BrightnessDown:
		brightnessDownRepeater.Stop()

	case conf.Combinations.VolumeUp:
		volumeUpRepeater.Stop()

	case conf.Combinations.VolumeDown:
		volumeDownRepeater.Stop()

	default:
		return eventprocessor.HANDLER_OK
	}

	return eventprocessor.HANDLER_ABORT
}

func fallbackProcessor(event *evdev.InputEvent) int {
	/*
		if event.Code != 0 {
			fmt.Println(event)
		}
	*/

	return eventprocessor.HANDLER_OK
}
