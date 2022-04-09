package repository

import (
	"context"
	"uahSalaryBot/external"
	"uahSalaryBot/infrastructure/domain"
)

type user struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Username  string
}

type User struct {
	client *external.DbClient
	model  *user
}

func NewUserRepository(c *external.DbClient) *User {
	return &User{client: c}
}

func (u *User) Create(_ context.Context) error {
	return u.client.Create(u.model).Error
}

func (u *User) ReadOne(ctx context.Context) error {
	return nil
}

func (u *User) Update(ctx context.Context) error {
	return nil
}

func (u *User) Delete(ctx context.Context) error {
	return nil
}

//FindOrCreate find by Username or create if not exist
//bool parameter returns true if user is found, false if created
func (u *User) FindOrCreate(_ context.Context, d *domain.User) error {
	u.setModel(d)
	db := u.client.FirstOrCreate(u.model, &user{Username: u.model.Username})

	if db.Error != nil {
		return db.Error
	}

	return nil
}

//setModel convert domain model to repository model. Must be used when run methods from BaseUser interface
func (u *User) setModel(domain *domain.User) *User {
	u.model = &user{
		ID:        domain.ID,
		FirstName: domain.FirstName,
		LastName:  domain.LastName,
		Username:  domain.Username,
	}

	return u
}
