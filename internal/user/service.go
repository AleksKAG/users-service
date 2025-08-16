package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(email string) (*User, error) {
	user := &User{Email: email}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetUserByID(id uint32) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdateUserByID(id uint32, user *User) (*User, error) {
	if err := s.repo.Update(id, user); err != nil {
		return nil, err
	}
	return s.repo.GetByID(id)
}

func (s *Service) DeleteUserByID(id uint32) error {
	return s.repo.Delete(id)
}

func (s *Service) GetAllUsers(page, pageSize uint32) ([]*User, error) {
	return s.repo.GetAll(page, pageSize)
}
