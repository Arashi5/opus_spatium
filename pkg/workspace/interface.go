package workspace

type Execute interface {
	Exec([]string) *error
}

type WorkSpace interface {
	GetImports()
	GetStreams() error
	GetLogger() error
	GetError() error
	GetGarbageCollection() error
	GetDraft()
}
