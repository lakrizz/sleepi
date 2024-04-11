package mpd

import (
	"context"
	"log"
	"os/exec"
)

// this package contains the listener
// which is not included in mpd due to the architecture of mpd itself
// (this is feature, not a bug! :D)

// as of now this will just be a mplayer running in a goroutine

type listener struct {
	running bool
	cancel  chan bool
	cmd     *exec.Cmd
}

func createListener(ctx context.Context, url string, cancel chan bool) *listener {
	l := &listener{
		running: false,
		cancel:  cancel,
		cmd:     exec.CommandContext(ctx, "mplayer", url),
	}

	return l
}

func (l *listener) run() {
	l.running = true
	log.Println("listener is now running")
	err := l.cmd.Start() // use builtin logic to prevent double execution
	if err != nil {
		return
	}

	l.cmd.Wait()
	log.Println("listener is not running anymore")
	l.running = false
}
