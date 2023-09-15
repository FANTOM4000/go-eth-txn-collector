package ports

type BlockQueHandler interface {
	ListenNewBlock()
}