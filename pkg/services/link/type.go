package links

type LinkService interface {
	Shorten(longLink string) (shortLink string, err error)
	GetOriginalFromShorten(shortLink string) (longLink string, err error)
}
