package audioplayer

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"krizz.org/sleepi/pkg/library"
)

type driver struct{}

func (d *driver) init() {
}

func (d *driver) load(file *library.File) error {

	data, err := file.Read()
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(io.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return err
	}

	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	log.Println(streamer, format)

	speaker.Play(streamer)
	select {}

	return nil
}
