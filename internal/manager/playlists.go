package manager

import "krizz.org/sleepi/pkg/playlist"

type PlaylistManager struct {
	Playlists []*playlist.Playlist
}

func GetPlaylistManager() (*PlaylistManager, error) {
	return &PlaylistManager{Playlists: make([]*playlist.Playlist, 0)}, nil
}

func (pm *PlaylistManager) AddPlaylist(playlist *playlist.Playlist) error {
	if !pm.isInList(playlist) {
		pm.Playlists = append(pm.Playlists, playlist)
	}
	return nil
}

func (pm *PlaylistManager) isInList(playlist *playlist.Playlist) bool {
	for _, v := range pm.Playlists {
		if v.Id == playlist.Id {
			return true
		}
	}
	return false
}
