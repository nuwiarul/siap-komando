package domain

import "gopkg.in/guregu/null.v4"

type User struct {
	ID        int
	Username  string
	Password  string
	Role      string
	CreatedAt null.Time
	UpdatedAt null.Time
}
