package main

import (
	"github.com/google/uuid"
	"github.com/k0kubun/pp"

	"github.com/lakrizz/sleepi/pkg/library"
	"github.com/lakrizz/sleepi/pkg/player/mpd"
)

func main() {

	bp, err := mpd.NewMPDPlayer()
	if err != nil {
		pp.Println(err)
		return
	}

	err = bp.Queue(&library.File{
		Path: "./cmd/sandbox/ttng.mp3",
		Id:   uuid.New(),
	})

	if err != nil {
		panic(err)
	}

	pp.Println(bp.Play())
}
