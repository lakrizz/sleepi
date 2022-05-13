package audioplayer

import (
	"time"

	"github.com/k0kubun/pp"
)

func (a *Audioplayer) StartVolumeGradient(start, end int, d time.Duration) {
	if a.volume_gradient_running {
		a.chan_volume_stop <- struct{}{} // stops all former functions
		a.volume_gradient_running = false
	}

	go func(a *Audioplayer, start, target int, d time.Duration, stop chan struct{}) {
		a.volume_gradient_running = true
		end := time.Now().Add(d)
		for time.Now().Before(end) {
			select {
			case <-stop:
				return
			default:
				vol_perc := 100 - (time.Until(end).Seconds())/d.Seconds()*float64(100)
				vol := float64(target) * (vol_perc / 100)
				pp.Println(time.Until(end).Milliseconds(), vol_perc, vol)
				a.SetVolume(int(vol))
				time.Sleep(time.Millisecond * 100) // 10 updates per seconds ought to be enough
			}
			a.volume_gradient_running = false // done
		}

	}(a, start, end, d, a.chan_volume_stop)
}
