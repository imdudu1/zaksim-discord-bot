package db

import (
	"entgo.io/ent/examples/start/ent"
	"sync"
)

var m sync.Mutex
var instance *ent.Client

func Instance() (*ent.Client, error) {
	m.Lock()
	defer m.Unlock()
	if instance == nil {
		client, err := ent.Open("postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
		if err != nil {
			return nil, err
		}
		instance = client
	}
	return instance, nil
}
