package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(u User) (User, error) {
	result := r.db.Create(&u)
	return u, result.Error
}

func (r *Repository) Get(id uint32) (User, error) {
	var u User
	result := r.db.First(&u, id)
	return u, result.Error
}

func (r *Repository) Update(u User) (User, error) {
	result := r.db.Model(&User{}).Where("id = ?", u.ID).Update("email", u.Email)
	if result.Error != nil {
		return User{}, result.Error
	}
	return u, nil
}

func (r *Repository) Delete(id uint32) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *Repository) List(offset, limit int) ([]User, error) {
	var users []User
	result := r.db.Offset(offset).Limit(limit).Find(&users)
	return users, result.Error
}
