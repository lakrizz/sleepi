package main

import (
	"github.com/k0kubun/pp"
	"krizz.org/sleepi/pkg/audioplayer"
	"krizz.org/sleepi/pkg/library"
	"krizz.org/sleepi/pkg/playlist"
)

func main() {
	lib, err := library.GetLibrary()
	if err != nil {
		panic(err)
	}
	pp.Println(lib.Files)

	// test_file := "/home/krizz/mp3/stickauto/001_Grooveshark/Ryan Hemsworth - Snow In Newark Ft. Dawn Golden.mp3"
	// dat, err := ioutil.ReadFile(test_file)
	// if err != nil {
	// 	panic(err)
	// }

	// err = lib.AddFile(dat, filepath.Base(test_file))
	// if err != nil {
	// 	panic(err)
	// }

	pl, err := playlist.NewPlaylist("test")
	if err != nil {
		panic(err)
	}

	for _, v := range lib.GetAllFiles() {
		pl.Add(v)
	}

	rf, err := pl.GetRandomFile()
	if err != nil {
		panic(err)
	}

	audioplayer.Audioplayer.Add(rf)
	audioplayer.Audioplayer.Play()

	select {}
}
