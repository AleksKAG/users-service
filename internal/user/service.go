package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(email string) (User, error) {
	return s.repo.Create(User{Email: email})
}

func (s *Service) GetUser(id uint32) (User, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateUser(id uint32, email string) (User, error) {
	return s.repo.Update(User{ID: id, Email: email})
}

func (s *Service) DeleteUser(id uint32) error {
	return s.repo.Delete(id)
}

func (s *Service) ListUsers(page, pageSize int) ([]User, error) {
	offset := (page - 1) * pageSize
	return s.repo.List(offset, pageSize)
}
