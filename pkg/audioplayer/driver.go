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

func (d *driver) init(random bool) error {
	c, err := mpd.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		return err
	}
	c.Random(random)
	c.Single(false)
	c.Repeat(false)
	// c.Clear()
	d.client = c
	if err := c.Ping(); err != nil {
		return err
	}
	return nil
}

func (d *driver) check_connection() {
	if d.client.Ping() != nil {
		c, _ := mpd.Dial("tcp", "127.0.0.1:6600")
		d.client = c
	}
}

func (d *driver) add(file *library.File) error {
	d.check_connection()
	mpd_filename := fmt.Sprintf("file://%v", file.Path)
	log.Println("mpd filename:", mpd_filename)
	return d.client.Add(mpd_filename)
}

func (d *driver) play() error {
	d.check_connection()
	err := d.client.Play(0)
	return err
}

func (d *driver) stop() error {
	d.check_connection()
	err := d.client.Stop()
	return err
}

func (d *driver) setvolume(i int) error {
	d.check_connection()
	return d.client.SetVolume(i)
}

func (d *driver) update() error {
	_, err := d.client.Update("")
	return err
}
