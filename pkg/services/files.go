package services

import (
	"mime/multipart"
	"net/http"
)

func ParseFiles(r *http.Request) ([]*multipart.FileHeader, error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	attachments := make([]*multipart.FileHeader, 0)
	for _, v := range r.MultipartForm.File {
		if len(v) == 0 {
			continue
		}

		attachments = append(attachments, v...)
	}

	return attachments, nil
}
