package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amirhnajafiz/hls/internal/http/handler"
)

// configure the songs' directory name
const songsDir = "songs"

func New(cfg Config) {
	if _, err := os.Stat(songsDir); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			_ = os.Mkdir(songsDir, os.ModePerm)

			log.Println("Dir created")
		} else {
			log.Println("Dir exists")
		}
	}

	// handler init
	h := handler.Handler{}

	// root
	http.Handle("/", h.Home())
	// add a handler to play song files
	http.Handle("/play", h.AddHeaders(http.FileServer(http.Dir(songsDir))))
	// add a handler for uploading a file
	http.Handle("/upload", h.UploadFile(songsDir))

	log.Printf("Starting server on %v\n", cfg.Port)
	log.Printf("Serving %s on HTTP port: %v\n", songsDir, cfg.Port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil))
}
