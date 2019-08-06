package youtube

import (
	"errors"
	"net/http"

	"../../models"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var apikey string = "AIzaSyCb5A4o-1ncYKJ4DmADZJgvFMrYqs2i4jw"
var maxresults int64 = 1

type searcher struct {
	service     *youtube.Service
	initialized bool
	messages    chan *models.Message
}

func (s *searcher) init() error {
	client := &http.Client{
		Transport: &transport.APIKey{Key: apikey},
	}

	svc, err := youtube.New(client)

	if err != nil {
		return err
	}

	s.service = svc
	s.initialized = true

	return nil
}

func (s *searcher) SearchVideos(keyword string) ([]*models.Video, error) {
	if !s.initialized {
		return nil, errors.New("should initialize the searcher first, right?")
	}

	s.messages <- models.CreateMessage(keyword, models.MSG_SEARCH)
	call := s.service.Search.List("id,snippet").Q(keyword).MaxResults(maxresults)
	response, err := call.Do()
	s.messages <- models.CreateMessage(keyword, models.MSG_SEARCH_DONE)

	if err != nil {
		return nil, err
	}

	lst := make([]*models.Video, 0)

	for _, v := range response.Items {
		switch v.Id.Kind {
		case "youtube#video":
			lst = append(lst, &models.Video{Id: v.Id.VideoId, Title: v.Snippet.Title})
		}
	}

	return lst, nil
}

// TODO(@kk): add playlists prolly?
