package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetByID(id uint32) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Update(id uint32, user *User) error {
	return r.db.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

func (r *Repository) Delete(id uint32) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *Repository) GetAll(page, pageSize uint32) ([]*User, error) {
	var users []*User
	offset := (page - 1) * pageSize
	if err := r.db.Offset(int(offset)).Limit(int(pageSize)).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
