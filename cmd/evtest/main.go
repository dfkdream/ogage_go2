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
	"time"
)

func main() {
	fmt.Println("Evtest started. Exiting in 20 sec...")

	conf := config.Load("/etc/ogage/config.yml")

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

				fmt.Println(i, event)
			}
		}(i, inputDevice)
	}

	time.Sleep(20 * time.Second)
}
