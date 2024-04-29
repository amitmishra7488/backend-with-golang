// models/user_model.go
package models

import (
	"context"
	"fmt"
	"golang-backend/db"

	"github.com/uptrace/bun"
)

type User struct {
    bun.BaseModel `bun:"table:users"`
    UserId        int64  `bun:"id,pk,autoincrement" json:"user_id"`
    Username      string `bun:"username" json:"username"`
    Email         string `bun:"email" json:"email"`
    Password      string `bun:"password" json:"-"`
    Role          string `bun:"role,default:user" json:"role"`
}


// Constants for valid user roles
const (
	RoleAdmin  = "admin"
	RoleAuthor = "author"
	RoleUser   = "user"
)

// SetRole sets the role of the user, allowing only specific values.
func (u *User) SetRole(role string) error {
	switch role {
	case RoleAdmin, RoleAuthor, RoleUser:
		u.Role = role
	default:
		return fmt.Errorf("invalid role: %s", role)
	}
	return nil
}

type UserRepository struct{}

func (m UserRepository) Create(ctx context.Context, user *User) error {
	_, err := db.GetDB().NewInsert().Model(user).Exec(ctx)
	return err
}

func (m UserRepository) QueryAll(ctx context.Context) (data []*User, err error) {
	err = db.GetDB().NewSelect().Model(&data).Scan(ctx, &data)
	return data, err
}

func (m UserRepository) DeleteUserById(ctx context.Context, id int64) error {
	_, err := db.GetDB().NewDelete().Model((*User)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}

func (m UserRepository) UpdateUserById(ctx context.Context, id int64, newData *User) error {
	_, err := db.GetDB().NewUpdate().Model(newData).Where("id = ?", id).Exec(ctx)
	return err
}

func (m *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
    var user User
    err := db.GetDB().NewSelect().Model(&user).
        Where("email = ?", email).
        Scan(ctx)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (m *UserRepository) GetUserProfile(ctx context.Context, userId int64) (*User, error) {
    var user User
    err := db.GetDB().NewSelect().Model(&user).
        Where("id = ?", userId).
        Scan(ctx)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
