package ports

type BlockQueRepositories interface {
	Produce(number uint64) error
}
