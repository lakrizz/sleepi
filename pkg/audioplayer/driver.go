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

func (d *driver) init(random bool) {
	c, err := mpd.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		log.Println(err)
	}
	c.Random(random)
	c.Single(false)
	c.Repeat(false)
	c.Clear()
	d.client = c
}

func (d *driver) add(file *library.File) error {
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
