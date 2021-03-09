package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/aid95/zaksim-discord-bot/ent"
	_ "github.com/lib/pq"
)

var m sync.Mutex

// Connection DB 연결 정보
type Connection struct {
	Client *ent.Client
	Ctx    context.Context
}

var conn *Connection

// Instance DB 연결 개체를 반환받기 위한 함수
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

// Open DB 연결 작업 전 실질적인 연결을 위한 함수
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

// Close DB 작업 완료 후 자원을 반환하기 위한 함수
func Close() error {
	if conn != nil {
		if err := conn.Client.Close(); err != nil {
			return err
		}
	}
	return nil
}
