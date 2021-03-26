package keys

type KeyServiceImpl struct{}

func NewKeyService() KeyService {
	return &KeyServiceImpl{}
}
func (s *KeyServiceImpl) Get() (key string, err error) {
	return "", nil
}
