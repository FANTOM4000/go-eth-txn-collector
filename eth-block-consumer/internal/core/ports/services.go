package ports

type ConsumerService interface {
	Consume(blockHash string) error
}
