package main

import (
	"github.com/go-vgo/robotgo"
)

func moveMouse(xy []int) {
	robotgo.MoveRelative(xy[0], xy[1])
}

func typeText(text string) {
	robotgo.TypeStr(text)
}

func pressKey(key string) {
	robotgo.KeySleep = 100
	robotgo.KeyTap(key)
}

var keyMap map[string]string

func initKeymap() {
	keyMap = map[string]string{
		"volumeup":   "audio_vol_up",
		"volumedown": "audio_vol_down",
		"back":       "left",
		"forward":    "right",
		"playpause":  "space",
	}
}

func refreshKeymap() {
	if config.Volume == "default" {
		// media keys
		keyMap["volumeup"] = "audio_vol_up"
		keyMap["volumedown"] = "audio_vol_down"
	} else {
		// up/down
		keyMap["volumeup"] = "up"
		keyMap["volumedown"] = "down"
	}

	if config.Play == "default" {
		// spacebar
		keyMap["playpause"] = "space"
	} else {
		// media Key
		keyMap["playpause"] = "audio_play"
	}

	if config.Seek == "default" {
		// left/right
		keyMap["back"] = "left"
		keyMap["forward"] = "right"
	} else {
		// media keys
		keyMap["back"] = "audio_prev"
		keyMap["forward"] = "audio_next"
	}
}
