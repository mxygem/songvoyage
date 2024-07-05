package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Setlist represents data about a setlist including how long it should remain available.
// For temporary setlists, their expiry will be set to creation time plus the
// temporaryPlaylistLifespan. For persisted setlists, their expiry may not be set or may
// be set to a yet to be determined lifespan. (TODO)
type Setlist struct {
	ID     primitive.ObjectID
	Name   string
	Expiry time.Time
	Songs  []*Song
}

// Song represents data about a particular song
type Song struct {
	Name   string
	Artist string
}

func setlist(ctx context.Context, db finder, name []byte) (*Setlist, error) {
	fmt.Printf("name: %s\n", name)
	return nil, nil
}

func delete(ctx context.Context, db deleter, name []byte) error {
	return nil
}

func clear(ctx context.Context, db clearer, name []byte) error {
	return nil
}

// TODO: Potentially combine save and update
func save(ctx context.Context, db saver, name []byte) error {
	return nil
}

func update(ctx context.Context, db updater, name []byte) error {
	return nil
}

func addSong(ctx context.Context, db songer, name []byte) error {
	return nil
}

func removeSong(ctx context.Context, db songer, name []byte) error {
	return nil
}
