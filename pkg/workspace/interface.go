package workspace

type WorkSpace interface {
	GetImports()
	GetStreams() error
	GetLogger() error
	GetError() error
	GetGarbageCollection()
}
