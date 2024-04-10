package mpd

import (
	"fmt"
	"os"
	"path"

	"github.com/fhs/gompd/v2/mpd"
	"github.com/k0kubun/pp"

	"github.com/lakrizz/sleepi/pkg/library"
	"github.com/lakrizz/sleepi/pkg/player"
)

type MPDPlayer struct {
	mpdHost   string
	volume    int
	musicPath string
	listener  *listener
	client    *mpd.Client
}

type MPDOption func(*MPDPlayer) error

func NewMPDPlayer(opts ...MPDOption) (*MPDPlayer, error) {
	p := &MPDPlayer{
		volume:    100,
		mpdHost:   "127.0.0.1",
		musicPath: "./run/music",
	}

	for _, o := range opts {
		err := o(p)
		if err != nil {
			return nil, err
		}
	}

	if err := p.connect(); err != nil {
		return nil, err
	}

	p.SetVolume(p.volume)

	// create listener
	l := createListener(fmt.Sprintf("http://%s:8000", p.mpdHost))
	p.listener = l

	return p, nil
}

func (mp *MPDPlayer) connect() error {
	cl, err := mpd.Dial("tcp", fmt.Sprintf("%s:6600", mp.mpdHost))
	if err != nil {
		return err
	}
	mp.client = cl
	return nil
}

func (mp *MPDPlayer) Play() error {
	if !mp.listener.running {
		go mp.listener.run()
	}
	return mp.client.Play(-1)
}

func (mp *MPDPlayer) Stop() error {
	return mp.client.Stop()
}

func (mp *MPDPlayer) Pause() error {
	return mp.client.Pause(true)
}

func (mp *MPDPlayer) Skip() error {
	return mp.client.Next()
}

func (mp *MPDPlayer) SetVolume(volume int) error {
	mp.volume = min(100, max(0, volume)) // limit the volume to 100%
	return mp.client.SetVolume(mp.volume)
}

func (mp *MPDPlayer) GetVolume() (int, error) {
	return mp.volume, nil
}

// Queue method adds a song, updates the db and then queues the song
func (mp *MPDPlayer) Queue(f *library.File) error {
	if !f.Exists() {
		return player.ErrFileNotFound
	}
	targetFilename := path.Join(mp.musicPath, f.Name())

	if _, err := os.Stat(targetFilename); os.IsNotExist(err) {
		dat, err := f.Read()
		if err != nil {
			return err
		}

		err = os.WriteFile(targetFilename, dat, 0777)
		if err != nil {
			return err
		}

		err = mp.refresh()
		if err != nil {
			return err
		}
	}

	return mp.client.Add(f.Name())
}

func (mp *MPDPlayer) QueueMany(_ []*library.File, _ bool) error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) GetCurrentSong() (player.Song, error) {
	mpdAttrs, err := mp.client.CurrentSong()
	if err != nil {
		return player.NoSong, err
	}
	pp.Println(mpdAttrs)
	return player.NoSong, nil
}

func (mp *MPDPlayer) refresh() error {
	_, err := mp.client.Rescan("")
	return err
}
