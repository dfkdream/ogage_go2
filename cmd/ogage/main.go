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
	"ogage_go2/internal/eventprocessor"
	"os"
	"sync"
)

// TODO: Find safer way to store config
var conf *config.Config

func init() {
	f, err := os.OpenFile("/var/log/ogage.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		slog.Error(
			"Failed to open log file. Falling back to stdout.",
			"err", err,
		)

		return
	}

	logger := slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	slog.SetDefault(logger)
}

func main() {
	conf = config.Load("/etc/ogage/config.yml")

	p := eventprocessor.New()
	p.Register(powerButtonProcessor)
	// no need to process headphones here
	// as it is automatically handled
	//p.Register(headphoneProcessor)
	p.Register(joypadPressProcessor)
	p.Register(joypadReleaseProcessor)
	p.Register(fallbackProcessor)

	for i, inputDevice := range conf.InputDevices {
		go func(i int, device string) {
			dev, err := evdev.Open(device)
			if err != nil {
				slog.Error(
					"Failed to open device file.",
					"device", device,
					"err", err,
				)
				return
			}

			for {
				event, err := dev.ReadOne()
				if err != nil {
					slog.Error(
						"Failed to read event.",
						"device", device,
						"err", err,
					)
					continue
				}

				p.Process(event)
			}
		}(i, inputDevice)
	}

	// Wait forever
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
