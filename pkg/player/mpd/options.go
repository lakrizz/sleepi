package mpd

func WithAddress(address string) MPDOption {
	return func(m *MPDPlayer) error {
		m.mpdHost = address
		return nil
	}
}

func WithDefaultVolume(volume int) MPDOption {
	return func(m *MPDPlayer) error {
		m.volume = volume
		return nil
	}
}
