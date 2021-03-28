package algo

type AlgoService interface {
	GenerateKeys(name string) (keys []string, err error)
	EnsureAll() (err error)
}
