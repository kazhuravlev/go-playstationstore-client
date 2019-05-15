package client

type GameOverview struct {
	Name         string
	Link         string
	Price        string
	PricePsPlus  string
	ImagesSrcSet SrcSet
}

type ListingMeta struct {
	CurrentPage   int
	MaxPage       int
	MinPage       int
	ObjectsOnPage int
}
