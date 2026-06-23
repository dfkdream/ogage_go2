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
	f          func()
	delay      time.Duration
	interval   time.Duration
	delayTimer *time.Timer
	ticker     *time.Ticker
}

func New(f func()) *Repeater {
	//var r Repeater
	r := Repeater{
		f:        f,
		delay:    1 * time.Hour,
		interval: 1 * time.Hour,
	}

	ticker := time.NewTicker(r.interval)
	ticker.Stop()
	r.ticker = ticker

	delayTimer := time.AfterFunc(r.delay, func() {
		ticker.Reset(r.interval)
		f()
	})
	delayTimer.Stop()
	r.delayTimer = delayTimer

	go func() {
		for {
			<-ticker.C
			f()
		}
	}()

	return &r
}

func (r *Repeater) Start(delay, interval time.Duration) {
	r.delay = delay
	r.interval = interval
	r.delayTimer.Reset(delay)
	r.f()
}

func (r Repeater) Stop() {
	r.delayTimer.Stop()
	r.ticker.Stop()
}
