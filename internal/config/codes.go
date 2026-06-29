/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package config

import "ogage_go2/internal/evdev"

var eventCodeMap map[uint16]string = map[uint16]string{
	// D-Pad
	evdev.EVENT_DPAD_LEFT:  "LEFT",
	evdev.EVENT_DPAD_RIGHT: "RIGHT",
	evdev.EVENT_DPAD_UP:    "UP",
	evdev.EVENT_DPAD_DOWN:  "DOWN",

	// Triggers
	evdev.EVENT_TL:  "TL",
	evdev.EVENT_TR:  "TR",
	evdev.EVENT_TL2: "TL2",
	evdev.EVENT_TR2: "TR2",

	// Fn Buttons
	evdev.EVENT_TRIGGER_HAPPY1: "F1",
	evdev.EVENT_TRIGGER_HAPPY2: "F2",
	evdev.EVENT_TRIGGER_HAPPY3: "F3",
	evdev.EVENT_TRIGGER_HAPPY4: "F4",
	evdev.EVENT_TRIGGER_HAPPY5: "F5",
	evdev.EVENT_TRIGGER_HAPPY6: "F6",
}

const (
	VOLUME_UP       = "VOLUME_UP"
	VOLUME_DOWN     = "VOLUME_DOWN"
	BRIGHTNESS_UP   = "BRIGHTNESS_UP"
	BRIGHTNESS_DOWN = "BRIGHTNESS_DOWN"
	BRIGHTNESS_MIN  = "BRIGHTNESS_MIN"
	BRIGHTNESS_MAX  = "BRIGHTNESS_MAX"
	HOTKEY          = "HOTKEY"
)
