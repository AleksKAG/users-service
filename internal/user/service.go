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

func (s *Service) GetUser(id uint) (User, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateUser(id uint, email string) (User, error) {
	u, err := s.repo.Get(id)
	if err != nil {
		return User{}, err
	}
	u.Email = email
	return s.repo.Update(u)
}

func (s *Service) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) ListUsers(page, pageSize int) ([]User, error) {
	offset := (page - 1) * pageSize
	return s.repo.List(offset, pageSize)
}
