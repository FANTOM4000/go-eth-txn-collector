package ports

type ConsumerService interface {
	Consume(number uint64) error
}
