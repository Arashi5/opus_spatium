package draft

type Exec struct {
	repo repository
}

type repository struct{}

func NewRepo() *Exec {
	return &Exec{repo: repository{}}
}

func (e Exec) Exec(args []string) *error {
	switch args[0] {
	case "v":
		e.repo.checkGlobalVar()
	case "f":
		e.repo.ftpFileChecker()
	}
	return nil
}
