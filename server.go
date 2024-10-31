package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//go:embed static
var embeddedFS embed.FS

func runHttpServer() {
	serverRoot, err := fs.Sub(embeddedFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(serverRoot))
	http.Handle("/", fileServer)
	http.HandleFunc("/btn", btnPressHandler)

	fmt.Printf("Starting server at port 80\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

func btnPressHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	command := string(bytes)

	fmt.Fprintf(w, "Ok")
	fmt.Print(command)

	//find command in keymap and if found, execute it
	keycode, commandValid := keyMap[command]
	if commandValid {
		pressKey(keycode)
		fmt.Println("Ok")
	} else {
		fmt.Println("Err")
	}

}

func trackpadMoveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	command := string(bytes)

	fmt.Fprintf(w, "Ok")
	fmt.Print(command)

	trackpadMoveString := strings.Split(command, ",")

	trackpadMove := make([]int, len(trackpadMoveString))
	for i, v := range trackpadMoveString {
		trackpadMove[i], _ = strconv.Atoi(v)
	}
	if len(trackpadMove) == 2 {
		moveMouse(trackpadMove)
	} else {
		fmt.Println("Malformed request body")
	}
}

func keyboardInputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	command := string(bytes)

	fmt.Fprintf(w, "Ok")
	fmt.Print(command)

	keyboardInput := strings.TrimSpace(command)
	if len(keyboardInput) > 0 {
		pressKey(keyboardInput)
	}
}
