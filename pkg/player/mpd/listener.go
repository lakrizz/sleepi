package mpd

import (
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

func createListener(url string, cancel chan bool) *listener {
	l := &listener{
		running: false,
		cancel:  cancel,
		cmd:     exec.Command("mplayer", url),
	}

	return l
}

func (l *listener) run() {
	l.running = true
	err := l.cmd.Run() // use builtin logic to prevent double execution
	l.running = false
	if err != nil {
		return
	}
}
