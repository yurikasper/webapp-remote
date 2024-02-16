package main

import (
	"gioui.org/app"
)

func main() {
	//generate local IP QR code
	generateQR()

	//initialize keystroke generator
	initKbd()

	//Start WebApp server as subroutine
	go runHttpServer()

	//start GIO UI (main window)
	startUI()

	//load config from file
	loadConfig()

	app.Main()
}
