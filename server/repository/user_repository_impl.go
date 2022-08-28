package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/nuwiarul/siap-komando/helper"
	"github.com/nuwiarul/siap-komando/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sqlx.Tx, username string) (domain.User, error) {
	query := `SELECT id, username, password, role, created_at, updated_at 
			FROM tbl_users WHERE username=?`
	query = tx.Rebind(query)
	stmt, err := tx.Preparex(query)
	helper.PanicIfError(err)
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, username)
	helper.PanicIfError(err)
	defer rows.Close()
	item := domain.User{}

	if rows.Next() {
		var createdAt string
		var updatedAt string
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Role,
			&createdAt, &updatedAt)
		helper.PanicIfError(err)
		item.CreatedAt = helper.StringToNullTime(createdAt)
		item.UpdatedAt = helper.StringToNullTime(updatedAt)
		return item, nil
	} else {
		return item, errors.New("username not found")
	}

}

func (repo UserRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, item domain.User) domain.User {
	query := `INSERT INTO tbl_users (username, password, role, created_at, updated_at) 
			VALUES (?, ?, ?, ?, ?) 
			RETURNING id`
	query = tx.Rebind(query)
	stmt, err := tx.Preparex(query)
	helper.PanicIfError(err)
	defer stmt.Close()
	var id int
	err = stmt.QueryRowContext(ctx, item.Username, item.Password, item.Role,
		item.CreatedAt, item.UpdatedAt).Scan(&id)
	helper.PanicIfError(err)
	item.ID = int(id)
	return item
}
func (repo UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sqlx.Tx, item domain.User) domain.User {
	sql := "UPDATE tbl_users SET password=?, updated_at=? WHERE id=?"
	sql = tx.Rebind(sql)
	_, err := tx.ExecContext(ctx, sql, item.Password, item.UpdatedAt, item.ID)
	helper.PanicIfError(err)
	return item
}
func (repo UserRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, item domain.User) {
	sql := "DELETE FROM tbl_users WHERE id=?"
	sql = tx.Rebind(sql)
	_, err := tx.ExecContext(ctx, sql, item.ID)
	helper.PanicIfError(err)
}
