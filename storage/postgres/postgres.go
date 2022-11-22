package postgres

import (
	"fmt"
	"library/storage"
	"log"

	"github.com/jmoiron/sqlx"
)

type postgresImpl struct {
	db       *sqlx.DB
	userRepo *userRepoImpl
	bookRepo *bookRepoImpl
}

func NewPostgres(pgConnStr string) storage.StorageI {
	db, err := sqlx.Connect("postgres", pgConnStr)

	if err != nil {
		fmt.Print(" xatolik shotta ")
		log.Fatalln(err)
	}
	return &postgresImpl{
		db: db,
		userRepo: &userRepoImpl{
			db: db,
		},
		bookRepo: &bookRepoImpl{
			db: db,
		},
	}
}

func (c *postgresImpl) User() storage.UserRepoI {
	return *c.userRepo
}

func (c *postgresImpl) Book() storage.BookRepoI {
	return *c.bookRepo
}

func (c *postgresImpl) CloseDb() error {
	return c.db.Close()
}
