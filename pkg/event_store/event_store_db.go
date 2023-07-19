package eventstore

import (
	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
)

type EventStoreDB struct {
	Client *esdb.Client
}

func NewEventStoreDB(connectString string) (*EventStoreDB, error) {
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

	return &EventStoreDB{Client: client}, nil
}

func (e *EventStoreDB) Close() error {
	return e.Client.Close()
}
