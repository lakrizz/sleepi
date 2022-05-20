package effects

import "time"

type VolumeWarmup struct {
	Duration    *time.Duration
	StartVolume int
	EndVolume   int
}
