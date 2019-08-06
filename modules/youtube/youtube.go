package youtube

import (
	"../../models"
)

type YouTube struct {
	Downloader *downloader
	Search     *searcher
	Messages   chan *models.Message
}

func CreateYouTubeWrapper() (*YouTube, error) {
	yt := &YouTube{}
	yt.Messages = make(chan *models.Message, 8)

	yt.Search = &searcher{initialized: false, messages: yt.Messages}
	err := yt.Search.init()

	if err != nil {
		return nil, err
	}

	yt.Downloader = &downloader{messages: yt.Messages}
	return yt, nil
}
