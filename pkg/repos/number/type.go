package number

type NumberRepo interface {
	Increment(numberId int, value int) (newValue int, err error)
}
