package consumer

type ConsumerHandler interface {
	ProcessOhlc(in []byte)
}
