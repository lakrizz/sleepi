package alarm

import (
	"krizz.org/sleepi/pkg/audioplayer"
	"krizz.org/sleepi/pkg/playlist"
)

func StartMPDPlaylist(playlist *playlist.Playlist, random bool) {
	audioplayer.Audioplayer.AddRange(playlist.Files)
	audioplayer.Audioplayer.Play(random)
}
