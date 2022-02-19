package alarm

import (
	"log"

	"krizz.org/sleepi/pkg/audioplayer"
	"krizz.org/sleepi/pkg/playlist"
)

func StartMPDPlaylist(playlist *playlist.Playlist, random bool) {
	audioplayer.Audioplayer.Clear()
	err := audioplayer.Audioplayer.AddRange(playlist.Files, random)
	if err != nil {
		log.Println(err)
	}
	audioplayer.Audioplayer.Play()
}
