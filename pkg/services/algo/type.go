package algo

type AlgoService interface {
	GenerateKeys() (values []string, err error)
	GetId() (algoId int)
}
type AlgoFactory interface {
	Get(algoName string) (AlgoService, error)
}
