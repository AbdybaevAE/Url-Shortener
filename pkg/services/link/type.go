package links

type LinkService interface {
	ShortenLink(link string) (key string, err error)
	GetLink(key string) (link string, err error)
}
