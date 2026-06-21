/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package evdev

// Button event values
const (
	VALUE_RELEASED = iota
	VALUE_PRESSED
)

// Audio event values
const (
	VALUE_INSERTED = iota
	VALUE_EJECTED
)

// Power
const (
	EVENT_POWER = 116
)

// Audio
const (
	EVENT_HEADPHONE_INSERT = 2
)

// Joypad
const (
	EVENT_TL = 310 + iota
	EVENT_TR
	EVENT_TL2
	EVENT_TR2
)

const (
	EVENT_DPAD_UP = 544 + iota
	EVENT_DPAD_DOWN
	EVENT_DPAD_LEFT
	EVENT_DPAD_RIGHT
)

const (
	EVENT_TRIGGER_HAPPY1 = 704 + iota
	EVENT_TRIGGER_HAPPY2
	EVENT_TRIGGER_HAPPY3
	EVENT_TRIGGER_HAPPY4
	EVENT_TRIGGER_HAPPY5
	EVENT_TRIGGER_HAPPY6
)
