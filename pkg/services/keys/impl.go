package keys

type KeyServiceImpl struct{}

func NewService() KeyService {
	return &KeyServiceImpl{}
}
func (s *KeyServiceImpl) Get() (string, error) {

	return "", nil
}
