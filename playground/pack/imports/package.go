package imports

import (
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
)

func (repository) simpleImportModule() {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}
