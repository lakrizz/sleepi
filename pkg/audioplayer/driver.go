package audioplayer

import (
	"fmt"
	"log"

	"github.com/fhs/gompd/v2/mpd"
	"krizz.org/sleepi/pkg/library"
)

type driver struct {
	client *mpd.Client
}

func (d *driver) init() {
	c, err := mpd.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		log.Println(err)
	}
	d.client = c
}

func (d *driver) load(file *library.File) error {
	err := d.client.Clear()
	if err != nil {
		return err
	}
	return d.client.Add(fmt.Sprintf("file://%v", file.Location))
}

func (d *driver) play() error {
	err := d.client.Play(0)
	return err
}

func (d *driver) stop() error {
	err := d.client.Stop()
	return err
}
