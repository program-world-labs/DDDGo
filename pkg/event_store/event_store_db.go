package eventstore

import (
	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
)

type StoreDB struct {
	Client *esdb.Client
}

func NewEventStoreDB(connectString string) (*StoreDB, error) {
	// region createClient
	settings, err := esdb.ParseConnectionString(connectString)
	if err != nil {
		return nil, err
	}

	// Creates a new Client instance.
	client, err := esdb.NewClient(settings)
	if err != nil {
		return nil, err
	}

	return &StoreDB{Client: client}, nil
}

func (e *StoreDB) Close() error {
	return e.Client.Close()
}
