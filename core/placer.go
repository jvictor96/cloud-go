package core

type Placer interface {
	PlaceArt(artWorks []ArtWork, terminal Terminal) []Placing
}
