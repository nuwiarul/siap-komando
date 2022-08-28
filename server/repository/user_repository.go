package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nuwiarul/siap-komando/model/domain"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, tx *sqlx.Tx, username string) (domain.User, error)
	Create(ctx context.Context, tx *sqlx.Tx, item domain.User) domain.User
	UpdatePassword(ctx context.Context, tx *sqlx.Tx, item domain.User) domain.User
	Delete(ctx context.Context, tx *sqlx.Tx, item domain.User)
}
