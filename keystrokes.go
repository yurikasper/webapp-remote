package main

import (
	"time"

	"github.com/micmonay/keybd_event"
)

var kb keybd_event.KeyBonding
var keyMap map[string]int

func initKbd() {
	//init keymap
	keyMap = make(map[string]int)
	keyMap["volumeup"] = keybd_event.VK_VOLUME_UP
	keyMap["volumedown"] = keybd_event.VK_VOLUME_DOWN
	keyMap["back"] = keybd_event.VK_LEFT
	keyMap["forward"] = keybd_event.VK_RIGHT
	keyMap["playpause"] = keybd_event.VK_SPACE

	//init keybd_event
	var err error
	kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
}

func pressKey(keyCode int) {
	// Select keys to be pressed
	kb.SetKeys(keyCode)

	//Press keys for 10 ms
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
}

func refreshKeymap() {
	if config.Volume == "default" {
		// media keys
		keyMap["volumeup"] = keybd_event.VK_VOLUME_UP
		keyMap["volumedown"] = keybd_event.VK_VOLUME_DOWN
	} else {
		// up/down
		keyMap["volumeup"] = keybd_event.VK_UP
		keyMap["volumedown"] = keybd_event.VK_DOWN
	}

	if config.Play == "default" {
		// spacebar
		keyMap["playpause"] = keybd_event.VK_SPACE
	} else {
		// media Key
		keyMap["playpause"] = keybd_event.VK_MEDIA_PLAY_PAUSE
	}

	if config.Seek == "default" {
		// left/right
		keyMap["back"] = keybd_event.VK_LEFT
		keyMap["forward"] = keybd_event.VK_RIGHT
	} else {
		// media keys
		keyMap["back"] = keybd_event.VK_MEDIA_PREV_TRACK
		keyMap["forward"] = keybd_event.VK_MEDIA_NEXT_TRACK
	}
}
