/*
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
*/

package config_test

import (
	"ogage_go2/internal/config"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	t.Cleanup(func() {
		err := os.RemoveAll("test_files")
		if err != nil {
			t.Fatal(err)
		}
	})

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    *config.Config
		wantErr bool
	}{
		{
			name: "load from file",
			path: "test_files/config.yml",
			want: &config.Config{
				InputDevices: []string{
					"/dev/input/event0",
					"/dev/input/event1",
					"/dev/input/event2",
				},
				BrightnessFile: config.DEFAULT_BRIGHTNESS_FILE,
				Power: config.Power{
					LongPressDuration: 1 * time.Second,
				},
				Command: config.Command{
					Delay:    500 * time.Millisecond,
					Interval: 80 * time.Millisecond,
				},
				JoypadBindings: map[string]string{
					"RIGHT": "VOLUME_UP",
					"LEFT":  "VOLUME_DOWN",
					"UP":    "BRIGHTNESS_UP",
					"DOWN":  "BRIGHTNESS_DOWN",
					"TL":    "VOLUME_DOWN",
					"TR":    "VOLUME_UP",
					"TL2":   "BRIGHTNESS_DOWN",
					"TR2":   "BRIGHTNESS_UP",
					"F5":    "HOTKEY",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := config.Load(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Load() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Load() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
