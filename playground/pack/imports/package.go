package imports

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

type Imports interface {
	SimpleImportModule()
}

type Repository struct {}

func NewRepo() *Repository  {
	return &Repository{}
}

func (Repository)SimpleImportModule()  {
	fmt.Println(simpleGitHub.AddTwo(5,6))
}
