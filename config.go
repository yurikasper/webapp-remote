package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

var config = struct {
	Volume string
	Play   string
	Seek   string
}{
	Volume: "default",
	Play:   "default",
	Seek:   "default",
}

func loadConfig() {
	//load gob file to struct
	readGob()

	//struct to parameters
	volumeSettingsGroup.Value = config.Volume
	playSettingsGroup.Value = config.Play
	seekSettingsGroup.Value = config.Seek

	//update keymap
	refreshKeymap()
}

func saveConfig() {
	//parameters to struct
	config.Volume = volumeSettingsGroup.Value
	config.Play = playSettingsGroup.Value
	config.Seek = seekSettingsGroup.Value

	//save struct to gob file
	writeGob()
}

func readGob() {
	file, _ := os.Open("config.gob")
	defer file.Close()
	decoder := gob.NewDecoder(file)
	decoder.Decode(&config)
	fmt.Println(config)
}

func writeGob() {
	file, _ := os.Create("config.gob")
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(config)
}
