package config

import (

	"example.com/mysqlgo/db"
	"example.com/mysqlgo/server"
)

type Config interface {
	GetServerCnf() server.Config
	GetDBCnf() db.Config
}

type config struct {
	svr server.Config
	dbC db.Config
}

func (c *config) GetServerCnf() server.Config {
	return c.svr
}

func (c *config) GetDBCnf() db.Config {
	return c.dbC
}

func NewConfig() Config {
	s := server.NewConfig()
	d := db.NewDBConfig()

	s.Print()
	d.Print()

	return &config{svr: s, dbC: d}
}
