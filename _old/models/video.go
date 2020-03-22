package models

import "fmt"

type Video struct {
	Id           string
	Title        string
	ThumbnailUrl string
}

func (v *Video) GetUrl() string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", v.Id)
}
