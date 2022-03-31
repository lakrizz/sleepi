package audioplayer

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"krizz.org/sleepi/pkg/library"
)

type audioplayer struct {
	driver *driver
}

var Audioplayer *audioplayer

func init() {
	Audioplayer = &audioplayer{driver: &driver{}}
	Audioplayer.driver.init(true)
}

func (a *audioplayer) Play() error {
	a.driver.play()

	cur, _ := a.driver.client.CurrentSong()
	log.Println("playing", cur)
	return nil
}

func (a *audioplayer) Stop() error {
	a.driver.stop()

	return nil
}

func (a *audioplayer) Add(file *library.File) error {
	// TODO: consider skipping on duplicates or let them flow?
	log.Println("new song:", file.Path)
	return a.driver.add(file)
}

func (a *audioplayer) AddRange(files []*library.File, shuffle bool) error {
	if len(files) == 0 {
		return errors.New("why bother adding an empty list of files?")
	}

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })
	}

	for _, v := range files {
		err := a.Add(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *audioplayer) Next() error {
	return a.driver.client.Next()
}

func (a *audioplayer) Clear() error {
	return a.driver.client.Clear()
}
