package db

import (
	"context"
	"github.com/aid95/zaksim-discord-bot/ent"
	"sync"
)

var m sync.Mutex

type Connection struct {
	Client *ent.Client
	Ctx    context.Context
}

var conn *Connection

func Instance() *Connection {
	m.Lock()
	defer m.Unlock()
	if conn == nil {
		if err := Connect(); err != nil {
			return nil
		}
	}
	return conn
}

func Connect() error {
	if conn == nil {
		client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
		if err != nil {
			return err
		}
		conn = &Connection{Client: client, Ctx: context.Background()}
	}
	return nil
}

func Close() {
	if conn != nil {
		conn.Client.Close()
	}
}
