package main

import (
	"github.com/google/uuid"
	"github.com/k0kubun/pp"

	"github.com/lakrizz/sleepi/pkg/library"
	"github.com/lakrizz/sleepi/pkg/player/beeplayer"
)

func main() {

	bp, err := beeplayer.NewBeepPlayer()
	if err != nil {
		return
	}

	err = bp.Queue(&library.File{
		Path: "/Users/kristofkipp/src/krizz.org/sleepi/cmd/sandbox/ttng.mp3",
		Id:   uuid.New(),
	})

	if err != nil {
		panic(err)
	}

	pp.Println(bp.Play())
}
