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
	"ogage_go2/internal/repeater"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func execWithLog(name string, arg ...string) {
	err := exec.Command(name, arg...).Run()
	if err != nil {
		fmt.Printf("%s: %s", name, err)
	}
}

func getBrightness() int {
	// TODO: Need better way to handle this
	if conf.BrightnessFile == "" {
		conf.BrightnessFile = config.DEFAULT_BRIGHTNESS_FILE
	}

	b, err := os.ReadFile(conf.BrightnessFile)
	if err != nil {
		fmt.Println("getBrightness:", err)
		return 255
	}

	brightness, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		fmt.Println("getBrightness:", err)
		return 255
	}

	return brightness
}

func setBrightness(brightness int) {
	// TODO: Need better way to handle this
	if conf.BrightnessFile == "" {
		conf.BrightnessFile = config.DEFAULT_BRIGHTNESS_FILE
	}

	err := os.WriteFile(conf.BrightnessFile, []byte(strconv.Itoa(brightness)), 0644)
	if err != nil {
		fmt.Println("setBrightness:", err)
	}
}

func updateBrightness(direction int) {
	updated := getBrightness()
	if updated < 20 {
		updated += 1 * direction
	} else {
		updated += 5 * direction
	}

	setBrightness(min(255, max(1, updated)))
}

func brightnessUp() {
	updateBrightness(1)
}

var brightnessUpRepeater = repeater.New(brightnessUp)

func brightnessDown() {
	updateBrightness(-1)
}

var brightnessDownRepeater = repeater.New(brightnessDown)

func volumeUp() {
	execWithLog("amixer", "-q", "sset", "Playback", "1%+")
}

var volumeUpRepeater = repeater.New(volumeUp)

func volumeDown() {
	execWithLog("amixer", "-q", "sset", "Playback", "1%-")
}

var volumeDownRepeater = repeater.New(volumeDown)

func audioSpeaker() {
	execWithLog("amixer", "-q", "sset", "'Playback Path'", "SPK")
}

func audioHeadphone() {
	execWithLog("amixer", "-q", "sset", "'Playback Path'", "HP")
}

func power() {
	// need better way to handle these
	execWithLog("pause.sh")
}

func powerWithHotkey() {
	execWithLog("finish.sh")
}
