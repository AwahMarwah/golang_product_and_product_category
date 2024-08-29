package health

func (s *service) Check() (err error) {
	return s.healthRepo.Ping()
}
