package mpd

import (
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

func createListener(url string) *listener {
	l := &listener{
		running: false,
		cancel:  make(chan bool),
		cmd:     exec.Command("mplayer", url),
	}

	return l
}

func (l *listener) run() {
	err := l.cmd.Start() // use builtin logic to prevent double execution
	if err != nil {
		return
	}

	log.Printf("started mplayer process with pid: %v\n", l.cmd.Process.Pid)
	l.running = true

	// run a second
	l.cmd.Wait()
	l.running = false
}
