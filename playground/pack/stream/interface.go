package stream

type Streams interface {
	StreamIn()
	StreamOut()
	StreamError()
}
