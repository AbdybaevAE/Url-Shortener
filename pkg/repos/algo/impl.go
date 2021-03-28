package algo

type repo struct{}

func New() AlgoRepo {
	return &repo{}
}
