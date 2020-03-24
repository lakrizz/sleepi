package player

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Player struct {
	playlist       *Playlist
	queue          []string
	playing_index  int
	volume         *effects.Volume
	current_volume float64
	max_volume     float64
	silence_volume float64
	playback_done  chan bool
	IsPlaying      bool
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
		p := &Player{playing_index: 0, current_volume: normal, silence_volume: silence, max_volume: normal, IsPlaying: false}
		p.playback_done = make(chan bool, 1)
		player = p
	}
	return nil
}

func (p *Player) LoadPlaylist(plname string, shuffle bool) error {
	playlist, err := LoadPlaylist(plname)
	if err != nil {
		log.Println("loading playlist? more like", err.Error())
		return err
	}
	p.playlist = playlist
	p.queue = make([]string, p.playlist.Length())
	copy(p.queue, playlist.Files)

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(p.queue), func(i, j int) {
			p.queue[i], p.queue[j] = p.queue[j], p.queue[i]
		})
	}

	return nil
}

func (p *Player) Play() error {
	if p.IsPlaying {
		return errors.New("player is already playing")
	}

	currentfile := p.queue[p.playing_index]
	done := make(chan bool) // this is the channel that gets filled whenever a song is done playing
	log.Println("trying to load file:", currentfile)
	f, err := p.loadfile(currentfile)
	if err != nil { // most likely no more songs left, eh :D
		return err
	}
	log.Println("1")
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Println("uhm?")
		log.Fatal(err)
		return err
	}

	log.Println("2")
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}
	defer speaker.Close()

	log.Println("3")
	baseSamplerate := format.SampleRate

	p.IsPlaying = true
	go p.volumeChanger()

	log.Println("4")

	for ; len(p.queue) > p.playing_index; p.playing_index++ {
		// now let's check if there's songs in the playlist
		currentfile := p.queue[p.playing_index]
		log.Println(fmt.Sprintf("now playing: %s", currentfile))

		f, err := p.loadfile(currentfile)
		if err != nil { // this should not happen, unless the file does not exist anymore
			p.IsPlaying = false
			return err
		}
		streamer, format, err = mp3.Decode(f)
		p.volume = &effects.Volume{Streamer: streamer, Base: 2, Volume: p.current_volume} // streamer is actually a pipe, nifty

		// callback being the last element signaling everything else that we're done here, sheeeeesh
		speaker.Play(beep.Resample(4, baseSamplerate, format.SampleRate, beep.Seq(p.volume, beep.Callback(func() {
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
	p.IsPlaying = false
	return nil
}

func (p *Player) loadfile(filename string) (*os.File, error) {
	if _, e := os.Stat(filename); os.ErrNotExist == e {
		return nil, os.ErrNotExist
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (p *Player) volumeChanger() {
	for {
		select {
		case <-p.playback_done:
			return
		default:
			if p.volume != nil && p.current_volume < p.max_volume {
				p.current_volume += 0.01
				p.volume.Volume = p.current_volume
				log.Println("new volume: ", p.current_volume)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
	return
}
