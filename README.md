# WebApp Remote

A portable media remote with it's own server and webapp written in Go.

First prototyped in PHP in 2020 and later rewritten in Processing Java, I've decided to rewrite it one last time in Go, giving it greater performance and portability by eliminating the Java Runtime requirement.

## Features

- HTTP server for providing web page locally
- QR code for easily opening webapp on your phone
- Progressive WebApp with full screen and icon
- Configurable options for each button (automatically saved to local folder)
- Single file and portable (all files embedded in the executable)

## Download

You can get the compiled binary over at the [Releases page](https://github.com/yurikasper/webapp-remote/releases)

## Screenshots

<img src="screenshots/main.PNG?raw=true" alt="Main screen" width="460"/> <img src="screenshots/settings.PNG?raw=true" alt="Settings page" width="290"/>

<img src="screenshots/mobile.png?raw=true" alt="Mobile Web Page" width="250"/> <img src="screenshots/add_to_home.png?raw=true" alt="Add to Home Screen" width="250"/> <img src="screenshots/webapp.png?raw=true" alt="WebApp" width="250"/>

## Building

To build from source, make sure you've got **go** and **go-winres** (only needed for the desktop icon) and use:
```
go-winres make
go build -ldflags -H=windowsgui -o "Webapp Remote.exe"
```
