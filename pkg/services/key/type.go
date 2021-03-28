package keys

type KeyService interface {
	Get() (key string, err error)
}
