package test

import (
	"context"
	"testing"

	"github.com/nuwiarul/siap-komando/app"
	"github.com/nuwiarul/siap-komando/helper"
	"github.com/nuwiarul/siap-komando/model/domain"
	"github.com/nuwiarul/siap-komando/repository"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestPassword(t *testing.T) {
	password := helper.HashPassword("12345678")
	t.Logf(password)
}

func TestFindUsername(t *testing.T) {
	db := app.NewSqlite("../database.sqlite3")

	tx, _ := db.Beginx()

	userRepository := repository.NewUserRepository()
	user, err := userRepository.FindByUsername(context.Background(), tx, "admin")
	helper.PanicIfError(err)
	tx.Commit()

	assert.Equal(t, user.Role, "superuser")

}

func TestCheckPassword(t *testing.T) {
	db := app.NewSqlite("../database.sqlite3")

	tx, _ := db.Beginx()

	userRepository := repository.NewUserRepository()
	user, err := userRepository.FindByUsername(context.Background(), tx, "admin")
	helper.PanicIfError(err)
	tx.Commit()

	assert.Equal(t, user.Role, "superuser")
	assert.Equal(t, helper.CheckPasswordHash("12345678", user.Password), true)
	assert.Equal(t, helper.CheckPasswordHash("12345", user.Password), false)

}

func TestCreateUser(t *testing.T) {
	db := app.NewSqlite("../database.sqlite3")

	tx, _ := db.Beginx()

	userRepository := repository.NewUserRepository()
	createdAt := helper.GetCurrentDatetime()
	userRepository.Create(context.Background(), tx, domain.User{
		Username:  "admin2",
		Password:  helper.HashPassword("12345678"),
		Role:      "superuser",
		CreatedAt: null.TimeFrom(createdAt),
		UpdatedAt: null.TimeFrom(createdAt),
	})
	tx.Commit()

	//assert.Equal(t, user.ID, 2)
}

func TestUpdatedPassword(t *testing.T) {
	db := app.NewSqlite("../database.sqlite3")

	tx, _ := db.Beginx()

	userRepository := repository.NewUserRepository()
	updatedAt := helper.GetCurrentDatetime()
	user, _ := userRepository.FindByUsername(context.Background(), tx, "admin1")
	assert.Equal(t, helper.CheckPasswordHash("12345678", user.Password), true)
	user.Password = helper.HashPassword("87654321")
	user.UpdatedAt = null.TimeFrom(updatedAt)
	user = userRepository.UpdatePassword(context.Background(), tx, user)
	tx.Commit()
	assert.Equal(t, helper.CheckPasswordHash("12345678", user.Password), false)
	assert.Equal(t, helper.CheckPasswordHash("87654321", user.Password), true)
}
