package firestoredb

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	client *firestore.Client
}

func New(ctx context.Context, projectID string) (*Firestore, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &Firestore{client: client}, nil
}

func (f *Firestore) GetClient() *firestore.Client {
	return f.client
}

func (f *Firestore) Close() error {
	if f.client != nil {
		return f.client.Close()
	}

	return nil
}
