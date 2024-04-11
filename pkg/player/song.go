package player

import "github.com/lakrizz/sleepi/pkg/models"

type Song struct {
	Length int
	Pos    int
	File   *models.File
}
