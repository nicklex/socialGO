package config

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func NewDB() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable", // data source name, refer https://github.com/jackc/pgx
	}))
	if err != nil {
		return nil, err
	} else {
		fmt.Println("db connected")
	}

	db.AutoMigrate(&Accounts{})

	return db, nil
}

func NewUser(db *gorm.DB, name, email, password string) (*Accounts, error) {
	ctx := context.Background()
	user := &Accounts{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := gorm.G[Accounts](db).Create(ctx, user)
	if err != nil {
		fmt.Println("no.", err.Error())
	}
	return user, nil
}
