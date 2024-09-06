package registry

type Record interface {
	FromMap(any) error
}
