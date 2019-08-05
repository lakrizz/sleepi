package youtube

type YouTube struct {
	Downloader *downloader
	Search     *searcher
}

func CreateYouTubeWrapper() (*YouTube, error) {
	yt := &YouTube{}
	yt.Search = &searcher{initialized: false}
	err := yt.Search.init()
	if err != nil {
		return nil, err
	}

	return yt, nil
}
