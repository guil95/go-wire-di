package database

type Mysql struct {}

type DB interface {
	GetItem() string
}

func NewMysqlDB() DB {
	return &Mysql{}
}

func (bd *Mysql) GetItem() string {
	return ""
}