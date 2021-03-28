package key

type repo struct{}

func New() KeyRepo {
	return &repo{}
}
