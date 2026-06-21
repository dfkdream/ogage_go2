/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package repeater

import "time"

type Repeater struct {
	f           func()
	delay       time.Duration
	interval    time.Duration
	delayticker *time.Ticker
	ticker      *time.Ticker
}

func New(f func(), delay time.Duration, interval time.Duration) *Repeater {
	ticker := time.NewTicker(interval)
	ticker.Stop()

	delayticker := time.NewTicker(delay)
	delayticker.Stop()

	go func() {
		for {
			select {
			case <-delayticker.C:
				delayticker.Stop()
				ticker.Reset(interval)
				f()
			case <-ticker.C:
				f()
			}
		}
	}()

	return &Repeater{
		f:           f,
		delay:       delay,
		interval:    interval,
		delayticker: delayticker,
		ticker:      ticker,
	}
}

func (r Repeater) Start() {
	r.f()
	r.delayticker.Reset(r.delay)
}

func (r Repeater) Stop() {
	r.delayticker.Stop()
	r.ticker.Stop()
}
