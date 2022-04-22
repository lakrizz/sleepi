package audioplayer

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"krizz.org/sleepi/pkg/library"
)

type Audioplayer struct {
	driver *driver
}

func GetAudioplayer() (*Audioplayer, error) {
	audioplayer := &Audioplayer{driver: &driver{}}
	err := audioplayer.driver.init(true)
	if err != nil {
		return nil, err
	}
	return audioplayer, nil
}

func (a *Audioplayer) Play() error {
	a.driver.play()

	cur, _ := a.driver.client.CurrentSong()
	log.Println("playing", cur)
	return nil
}

func (a *Audioplayer) Stop() error {
	a.driver.stop()

	return nil
}

func (a *Audioplayer) Add(file *library.File) error {
	// TODO: consider skipping on duplicates or let them flow?
	log.Println("new song:", file.Path)
	return a.driver.add(file)
}

func (a *Audioplayer) AddRange(files []*library.File, shuffle bool) error {
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

func (a *Audioplayer) Next() error {
	return a.driver.client.Next()
}

func (a *Audioplayer) Clear() error {
	return a.driver.client.Clear()
}

func (a *Audioplayer) SetVolume(volume int) error {
	if volume < 0 {
		return errors.New("volume should not be below 0")
	}

	if volume > 100 {
		return errors.New("volume should not be above 100")
	}

	return a.driver.setvolume(volume)
}
