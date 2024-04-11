package beeplayer

import (
	"fmt"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"

	"github.com/lakrizz/sleepi/pkg/models"
	"github.com/lakrizz/sleepi/pkg/player"
)

type BeepPlayer struct {
	volume      int
	state       int
	currentSong *player.Song

	queue []*models.File

	command chan func() error
}

func NewBeepPlayer(opts ...player.PlayerOption) (*BeepPlayer, error) {
	player := &BeepPlayer{
		volume:  0,
		state:   player.StateStopped,
		queue:   make([]*models.File, 0),
		command: make(chan func() error, 0),
	}

	for _, o := range opts {
		o(player)
	}

	return player, nil
}

// Play starts the player, the songs are played in the order they
// were added
func (be *BeepPlayer) Play() error {
	if len(be.queue) == 0 {
		return player.ErrQueueEmpty
	}

	f := be.queue[0]
	streamer, err := be.load(f)
	if err != nil {
		return err
	}

	done := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	be.state = player.StatePlaying
	<-done
	streamer.Close()

	return nil
}

func (be *BeepPlayer) Stop() error {
	speaker.Suspend()
	return nil
}

func (be *BeepPlayer) Pause() error {
	panic("not implemented") // TODO: Implement
}

func (be *BeepPlayer) SetVolume(volume int) error {
	be.volume = min(100, max(0, volume))
	return nil
}

func (be *BeepPlayer) GetVolume() (int, error) {
	return be.volume, nil
}

func (be *BeepPlayer) Queue(file *models.File) error {
	// pushing the given file at the end of the queue
	if !file.Exists() {
		return player.ErrFileNotFound
	}

	be.queue = append(be.queue, file)
	return nil
}

func (be *BeepPlayer) QueueMany(files []*models.File, shuffle bool) error {
	if shuffle {
		fmt.Println("TODO: shuffling list")
	}

	for _, f := range files {
		err := be.Queue(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func (be *BeepPlayer) Skip() error {
	panic("not implemented") // TODO: Implement
}

func (be *BeepPlayer) GetCurrentSong() (player.Song, error) {
	if be.currentSong == nil {
		return player.NoSong, player.ErrNoSongLoaded
	}

	return *be.currentSong, nil // we are returning a copy to keep the current song immutable
}

func (be *BeepPlayer) load(f *models.File) (beep.StreamSeekCloser, error) {
	rc, err := f.GetReadCloser()
	if err != nil {
		return nil, err
	}

	streamer, format, err := mp3.Decode(rc)
	if err != nil {
		return nil, err
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	return streamer, nil
}
