/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package eventprocessor

import "ogage_go2/internal/evdev"

const (
	HANDLER_OK = iota
	HANDLER_ABORT
)

type EventHandler func(*evdev.InputEvent) int

type EventProcessor struct {
	handlers []EventHandler
}

func New() *EventProcessor {
	var p EventProcessor

	p.handlers = make([]EventHandler, 0)

	return &p
}

func (e *EventProcessor) Register(handler EventHandler) {
	// TODO: check race conditions
	e.handlers = append(e.handlers, handler)
}

func (e EventProcessor) Process(event *evdev.InputEvent) {
	for _, h := range e.handlers {
		result := h(event)

		if result == HANDLER_ABORT {
			return
		}
	}
}
