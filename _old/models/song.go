package models

import "time"

type Song struct {
	Length       int
	Artist       string
	Title        string
	DownloadDate *time.Time
}
