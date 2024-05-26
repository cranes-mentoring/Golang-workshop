package service

func (s *Service) CompleteTask(ID string) error {
	return s.Repo.CompleteTask(ID)
}
