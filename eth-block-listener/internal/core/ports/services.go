package ports

type BlockQueService interface {
	Produce(number uint64) error
}
