/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package repeater_test

import (
	"ogage_go2/internal/repeater"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	i := 0
	r := repeater.New(func() {
		t.Log("Function fired.")
		i++
	}, 5*time.Millisecond, 1*time.Millisecond)

	t.Run("10ms wait", func(t *testing.T) {
		r.Start()
		time.Sleep(10 * time.Millisecond)
		r.Stop()
		time.Sleep(2 * time.Millisecond) // To check if Stop() works.

		if i != 7 {
			t.Errorf("Expected i=7 but got %d", i)
			t.Fail()
		}
	})

	t.Run("2ms wait", func(t *testing.T) {
		i = 0
		r.Start()
		time.Sleep(2 * time.Millisecond)
		r.Stop()

		if i != 1 {
			t.Errorf("Expected i=1 but got %d", i)
			t.Fail()
		}
	})
}
