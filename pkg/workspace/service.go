package workspace

func NewService(cfg *Config) *service {
	return &service{
		Arguments:  getAdditionalArgs(cfg.Arguments),
		Repository: getPGRepoCollection(),
	}
}

func (s service) GetDraft() *error {
	if err := s.Repository.D.Exec(s.Arguments); err != nil {
		return err
	}

	return nil
}

func (s service) GetImports() *error {
	if err := s.Repository.Imp.Exec(s.Arguments); err != nil {
		return err
	}

	return nil
}

func (s service) GetStreams() *error {
	if err := s.Repository.Str.Exec(s.Arguments); err != nil {
		return err
	}

	return nil
}

func (s service) GetLogger() *error {
	if err := s.Repository.Log.Exec(s.Arguments); err != nil {
		return err
	}

	return nil
}

func (s service) GetError() *error {
	if err := s.Repository.Err.Exec(s.Arguments); err != nil {
		return err
	}

	return nil
}

//time go run main.go gc slice s|ms|mns|mst
func (s service) GetGarbageCollection() *error {
	if err := s.Repository.GC.Exec(s.Arguments); err != nil {
		return err
	}

	return nil
}

func getAdditionalArgs(args []string) Arguments {
	var aa Arguments

	if len(args) < 2 {
		return aa
	}

	for i := 0; i < len(args); i++ {
		if i < 2 {
			continue
		}
		aa = append(aa, args[i])
	}

	return aa
}
