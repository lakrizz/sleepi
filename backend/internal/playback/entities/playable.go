package entities

type Playable struct {
	URI   string // i think every player has some kind of format to identify a resource, so we'll use this uri
	Title string

	Source Source
}
