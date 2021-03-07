package db

import (
	"context"
	"fmt"
	"github.com/aid95/zaksim-discord-bot/ent"
	_ "github.com/lib/pq"
	"os"
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
		if err := Open(); err != nil {
			return nil
		}
	}
	return conn
}

func Open() error {
	if conn == nil {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		db := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, db, pass))
		if err != nil {
			return err
		}

		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			return err
		}

		conn = &Connection{Client: client, Ctx: ctx}
	}
	return nil
}

func Close() error {
	if conn != nil {
		if err := conn.Client.Close(); err != nil {
			return err
		}
	}
	return nil
}
