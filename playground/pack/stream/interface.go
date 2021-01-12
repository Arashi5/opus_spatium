package stream

type Streams interface {
	streamIn()
	streamOut()
	streamError()
}
