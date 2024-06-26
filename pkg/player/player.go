package player

import (
	"errors"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/google/uuid"
	"github.com/lakrizz/sleepi/pkg/library"
)

type Player struct {
	playing_queue  []uuid.UUID
	playing_index  int
	volume         *effects.Volume
	current_volume float64
	max_volume     float64
	silence_volume float64
	playback_done  chan bool
	stop_volume    chan bool
	IsPlaying      bool
	library        *library.Library
}

var player *Player

func GetPlayer() (*Player, error) {
	if player == nil {
		return nil, errors.New("player is not initialized")
	}
	return player, nil
}

func InitPlayer(silence, normal float64) error {
	if player == nil {
		p := &Player{playing_index: 0, current_volume: silence, silence_volume: silence, max_volume: normal, IsPlaying: false}
		p.playback_done = make(chan bool, 1)
		p.stop_volume = make(chan bool, 1)
		player = p
	}
	return nil
}

func (p *Player) QueueSongs(songs []uuid.UUID, shuffle bool) error {
	copy(p.playing_queue, songs)

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(p.playing_queue), func(i, j int) {
			p.playing_queue[i], p.playing_queue[j] = p.playing_queue[j], p.playing_queue[i]
		})
	}

	return nil
}

func (p *Player) Play() error {
	if p.IsPlaying {
		return errors.New("player is already playing")
	}
	var glob_err error

	go func() {
		currentfile := p.playing_queue[p.playing_index]
		done := make(chan bool) // this is the channel that gets filled whenever a song is done playing
		f, err := p.loadfile(currentfile)
		if err != nil { // most likely no more songs left, eh :D
			glob_err = err
			return
		}
		_, format, err := mp3.Decode(f)
		if err != nil {
			glob_err = err
			return
		}

		baseSamplerate := format.SampleRate
		err = speaker.Init(baseSamplerate, baseSamplerate.N(time.Second/10))
		if err != nil {
			panic(err)
		}

		p.IsPlaying = true
		go p.volumeChanger()

		for ; len(p.playing_queue) > p.playing_index; p.playing_index++ {
			// now let's check if there's songs in the playlist
			currentfile = p.playing_queue[p.playing_index]

			f, err = p.loadfile(currentfile)
			if err != nil { // this should not happen, unless the file does not exist anymore
				p.IsPlaying = false
				glob_err = err
				return
			}
			streamer, _, err := mp3.Decode(f)
			if err != nil {
				panic(err)
			}

			p.volume = &effects.Volume{Streamer: streamer, Base: 2, Volume: p.current_volume} // streamer is actually a pipe, nifty

			// callback being the last element signaling everything else that we're done here, sheeeeesh
			speaker.Play(beep.Resample(2, baseSamplerate, format.SampleRate, beep.Seq(p.volume, beep.Callback(func() {
				done <- true
			}))))

			<-done // block until the song is played

			err = streamer.Close()
			if err != nil {
				p.IsPlaying = false
				panic(err)
			}
		}

		p.playback_done <- true
		p.stop_volume <- true
	}()
	<-p.playback_done
	p.IsPlaying = false
	return glob_err
}

func (p *Player) loadfile(id uuid.UUID) (*os.File, error) {
	file := p.library.GetFile(id)
	f, err := os.Open(path.Join(file.Path, file.Filename))
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (p *Player) volumeChanger() {
	for {
		select {
		case <-p.stop_volume:
			return
		default:
			if p.volume != nil && p.current_volume < p.max_volume {
				p.current_volume += 0.01
				p.volume.Volume = p.current_volume
				time.Sleep(50 * time.Millisecond)
			} else if p.current_volume == p.max_volume {
				// we don't need to adjust the volume anymore, stop this goroutine
				return
			}
		}
	}
}
