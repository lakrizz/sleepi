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

	// 	test_files := []string{
	// 		"/home/krizz/mp3/stickauto/12 - Herbst 14/sleeping at last - households-4SjHY54BwUc.mp3",
	// 		"/home/krizz/mp3/stickauto/12 - Herbst 14/Hazelton - Justin Vernon-mU6OQRzsQ5A.mp3",
	// 	}
	// 	for _, f := range test_files {
	// 		dat, err := ioutil.ReadFile(f)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		err = lib.AddFile(dat, filepath.Base(f))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}

	pl, err := playlist.NewPlaylist("test")
	if err != nil {
		panic(err)
	}

	for _, v := range lib.GetAllFiles() {
		pl.Add(v)
	}

	audioplayer.Audioplayer.AddRange(lib.GetAllFiles())

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
