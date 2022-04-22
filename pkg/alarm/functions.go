package alarm

import (
	"log"

	"krizz.org/sleepi/pkg/audioplayer"
	"krizz.org/sleepi/pkg/playlist"
)

func StartMPDPlaylist(player *audioplayer.Audioplayer, playlist *playlist.Playlist, random bool) {
	player.Clear()
	err := player.AddRange(playlist.Files, random)
	if err != nil {
		log.Println(err)
	}
	player.Play()
}
