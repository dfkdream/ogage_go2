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
	"ogage_go2/internal/eventprocessor"
	"sync/atomic"
	"time"
)

var hotkeyPressed atomic.Bool

var powerButtonTimer *time.Timer

func powerButtonProcessor(event *evdev.InputEvent) int {
	if event.Code != evdev.EVENT_POWER {
		return eventprocessor.HANDLER_OK
	}

	if hotkeyPressed.Load() {
		powerWithHotkey()
		return eventprocessor.HANDLER_ABORT
	}

	conf := config.Get()

	if powerButtonTimer == nil {
		powerButtonTimer = time.AfterFunc(
			conf.Power.LongPressDuration, powerWithHotkey)
		powerButtonTimer.Stop()
	}

	if event.Value == evdev.VALUE_PRESSED {
		powerButtonTimer.Reset(conf.Power.LongPressDuration)
	} else {
		powerButtonTimer.Stop()
		power()
	}

	return eventprocessor.HANDLER_ABORT
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

func joypadPressProcessor(event *evdev.InputEvent) int {
	if event.Value != evdev.VALUE_PRESSED {
		return eventprocessor.HANDLER_OK
	}

	conf := config.Get()

	binding := conf.JoypadBinding(event.Code)
	if binding == config.HOTKEY {
		hotkeyPressed.Store(true)
		return eventprocessor.HANDLER_ABORT
	}

	if !hotkeyPressed.Load() {
		return eventprocessor.HANDLER_OK
	}

	switch binding {
	case config.BRIGHTNESS_UP:
		brightnessUpRepeater.Start(
			conf.Command.Delay, conf.Command.Interval)

	case config.BRIGHTNESS_DOWN:
		brightnessDownRepeater.Start(
			conf.Command.Delay, conf.Command.Interval)

	case config.BRIGHTNESS_MIN:
		brightnessMin()

	case config.BRIGHTNESS_MAX:
		brightnessMax()

	case config.VOLUME_UP:
		volumeUpRepeater.Start(
			conf.Command.Delay, conf.Command.Interval)

	case config.VOLUME_DOWN:
		volumeDownRepeater.Start(
			conf.Command.Delay, conf.Command.Interval)

	default:
		return eventprocessor.HANDLER_OK
	}

	return eventprocessor.HANDLER_ABORT
}

func joypadReleaseProcessor(event *evdev.InputEvent) int {
	if event.Value != evdev.VALUE_RELEASED {
		return eventprocessor.HANDLER_OK
	}

	binding := config.Get().JoypadBinding(event.Code)
	if binding == config.HOTKEY {
		hotkeyPressed.Store(false)

		brightnessUpRepeater.Stop()
		brightnessDownRepeater.Stop()
		volumeUpRepeater.Stop()
		volumeDownRepeater.Stop()

		return eventprocessor.HANDLER_ABORT
	}

	if !hotkeyPressed.Load() {
		return eventprocessor.HANDLER_OK
	}

	switch binding {
	case config.BRIGHTNESS_UP:
		brightnessUpRepeater.Stop()

	case config.BRIGHTNESS_DOWN:
		brightnessDownRepeater.Stop()

	case config.VOLUME_UP:
		volumeUpRepeater.Stop()

	case config.VOLUME_DOWN:
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
