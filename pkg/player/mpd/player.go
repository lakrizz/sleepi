package mpd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/fhs/gompd/v2/mpd"
	"github.com/k0kubun/pp"

	"github.com/lakrizz/sleepi/pkg/models"
	"github.com/lakrizz/sleepi/pkg/player"
)

type MPDPlayer struct {
	mpdHost   string
	volume    int
	musicPath string
	listener  *listener
	client    *mpd.Client

	cancel chan bool
}

type MPDOption func(*MPDPlayer) error

func NewMPDPlayer(ctx context.Context, opts ...MPDOption) (*MPDPlayer, error) {
	p := &MPDPlayer{
		volume:    100,
		mpdHost:   "127.0.0.1",
		musicPath: "./run/music",
		cancel:    make(chan bool),
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
	l := createListener(ctx, fmt.Sprintf("http://%s:8000", p.mpdHost), p.cancel)
	p.listener = l
	go p.listener.run()

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
func (mp *MPDPlayer) Queue(f *models.File) error {
	if !f.Exists() {
		return player.ErrFileNotFound
	}
	targetFilename := path.Join(mp.musicPath, f.Name())
	log.Printf("copying %v to %v", f.Path, targetFilename)

	if _, err := os.Stat(targetFilename); os.IsNotExist(err) {
		dat, err := f.Read()
		if err != nil {
			return fmt.Errorf("error while queueing file: %w", err)
		}

		err = os.WriteFile(targetFilename, dat, 0777)
		if err != nil {
			return fmt.Errorf("error while writing queued file: %w", err)
		}

		err = mp.refresh()
		if err != nil {
			return fmt.Errorf("error while refreshing mpd database: %w", err)
		}
	}

	log.Println("trying to play", f.Name())
	return mp.client.Add(f.Name())
}

func (mp *MPDPlayer) QueueMany(_ []*models.File, _ bool) error {
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
	_, err := mp.client.Update("")
	return err
}

func (mp *MPDPlayer) Clear() error {
	return mp.client.Clear()
}
