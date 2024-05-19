package repository

import (
	"github.com/singland1997/e-com/src/entity"
)

type user struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Roles    string `db:"roles"`
	CreateAt string `db:"create_at"`
	UpdateAt string `db:"update_at"`
}

func UserRepository() entity.User {
	return entity.User{}
}
