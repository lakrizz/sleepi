package player

import "github.com/lakrizz/sleepi/pkg/library"

type Song struct {
	Length int
	Pos    int
	File   *library.File
}
