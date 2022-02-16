package audioplayer

import (
	"errors"
	"log"

	"krizz.org/sleepi/pkg/library"
)

type audioplayer struct {
	driver *driver
	queue  []*library.File
}

var Audioplayer *audioplayer

func init() {
	Audioplayer = &audioplayer{queue: make([]*library.File, 0), driver: &driver{}}
	Audioplayer.driver.init()
}

func (a *audioplayer) Play() error {
	if len(a.queue) < 1 {
		return errors.New("queue is empty, can't play")
	}

	// TODO: driver needs to be called here with a.Queue[0] if len(queue)>0
	// dequeue item
	item := a.queue[0]
	a.queue = a.queue[1:]
	err := a.driver.load(item)
	if err != nil {
		return err
	}
	a.driver.play()

	log.Println("playing", item.Location)
	return nil
}

func (a *audioplayer) Stop() error {
	a.driver.stop()

	return nil
}

func (a *audioplayer) Add(file *library.File) error {
	// TODO: consider skipping on duplicates or let them flow?
	log.Println("new song:", file.Location)
	a.queue = append(a.queue, file)

	return nil
}

func (a *audioplayer) AddRange(files []*library.File) error {
	if len(files) == 0 {
		return errors.New("why bother adding an empty list of files?")
	}

	for _, v := range files {
		err := a.Add(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *audioplayer) GetQueue() []*library.File {
	return a.queue
}

func (a *audioplayer) Next() error {
	if len(a.queue) < 1 {
		return errors.New("no next song")
	}
	a.Stop()
	return a.Play()
}
