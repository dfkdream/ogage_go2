/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package main

import (
	"fmt"
	"ogage_go2/internal/config"
	"ogage_go2/internal/evdev"
	"ogage_go2/internal/eventprocessor"
	"sync"
)

// TODO: Find safer way to store config
var conf *config.Config

func main() {
	conf, _ = config.Load("dummy")

	p := eventprocessor.New()
	p.Register(powerButtonProcessor)
	// no need to process headphones here
	// as it is automatically handled
	//p.Register(headphoneProcessor)
	p.Register(hotkeyProcessor)
	p.Register(combinationPressProcessor)
	p.Register(combinationReleaseProcessor)
	p.Register(fallbackProcessor)

	for i, inputDevice := range conf.InputDevices {
		go func(i int, device string) {
			dev, err := evdev.Open(device)
			if err != nil {
				fmt.Println(err)
				return
			}

			for {
				event, err := dev.ReadOne()
				if err != nil {
					fmt.Println(err)
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
