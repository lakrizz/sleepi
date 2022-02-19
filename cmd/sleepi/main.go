package main

import (
	"log"
	"time"

	"krizz.org/sleepi/pkg/audioplayer"
	"krizz.org/sleepi/pkg/library"
	"krizz.org/sleepi/pkg/playlist"
	"krizz.org/sleepi/web"
)

func main() {

	web.Serve()

	lib, err := library.GetLibrary()
	if err != nil {
		panic(err)
	}

	pl, err := playlist.NewPlaylist("test")
	if err != nil {
		panic(err)
	}

	for _, v := range lib.GetAllFiles() {
		pl.Add(v)
	}

	audioplayer.Audioplayer.AddRange(lib.GetAllFiles(), false)

	err = audioplayer.Audioplayer.Play()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	log.Println("stopping now")
	err = audioplayer.Audioplayer.Next()
	if err != nil {
		log.Println(err)
	}

	for {
	}
}
