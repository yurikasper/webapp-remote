package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
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
		fmt.Println(" ...Ok")
	} else {
		fmt.Println(" ...Err")
	}

}
