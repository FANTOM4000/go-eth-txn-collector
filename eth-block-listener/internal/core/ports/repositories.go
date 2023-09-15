package ports

type BlockQueRepositories interface {
	Produce(hex string) error
}
