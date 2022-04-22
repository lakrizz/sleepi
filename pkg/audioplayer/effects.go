package audioplayer

import (
	"time"

	"github.com/k0kubun/pp"
)

func (a *Audioplayer) AddVolumeGradient(start, end int, d time.Duration) error {
	go func(a *Audioplayer, start, target int, d time.Duration) {
		end := time.Now().Add(d)
		for time.Now().Before(end) {
			vol_perc := 100 - (time.Until(end).Seconds())/d.Seconds()*float64(100)
			vol := float64(target) * (vol_perc / 100)
			pp.Println(time.Until(end).Milliseconds(), vol_perc, vol)
			a.SetVolume(int(vol))
			time.Sleep(time.Millisecond * 100) // 10 updates per seconds ought to be enough
		}

	}(a, start, end, d)
	return nil
}
