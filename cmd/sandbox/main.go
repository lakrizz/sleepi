package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"

	"github.com/lakrizz/sleepi/pkg/library"
	"github.com/lakrizz/sleepi/pkg/player/mpd"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bp, err := mpd.NewMPDPlayer(ctx)
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
	select {}
}
