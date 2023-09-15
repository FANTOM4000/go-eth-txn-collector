package ports

type BlockQueService interface {
	Produce(hex string) error
}
