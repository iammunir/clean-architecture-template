package repository

import (
	"gorm.io/gorm"
)

type Repository interface {
	Ping() error
}

type repository struct {
	mysqlConn *gorm.DB
}

func NewRepository(mysqlConn *gorm.DB) Repository {
	return &repository{
		mysqlConn: mysqlConn,
	}
}

func (repo *repository) Ping() error {
	db, err := repo.mysqlConn.DB()
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
