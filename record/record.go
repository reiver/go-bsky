package record

type Record interface {
	FromMap(any) error
}
