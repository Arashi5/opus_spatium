package imports

import (
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
)

type Exec struct {
	repo repository
}

type repository struct{}

func NewRepo() *Exec {
	return &Exec{repo: repository{}}
}

func (e Exec) Exec(_ []string) *error {
	e.repo.simpleImportModule()
	return nil
}

func (repository) simpleImportModule() {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}
