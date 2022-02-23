package repository

import "github.com/guil95/go-wire-di/pkg/database"

type MysqlRepository struct {
	bd database.DB
}

func NewMysqlRepository(bd database.DB) Repository {
	return &MysqlRepository{
		bd: bd,
	}
}

func (repo *MysqlRepository) Get() string {
	return ""
}