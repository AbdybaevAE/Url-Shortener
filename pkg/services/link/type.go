package links

type LinkService interface {
	ShortenLink(link string) (linkKey string, err error)
	GetLink(linkKey string) (link string, err error)
	VisitByKey(key string) (err error)
}
