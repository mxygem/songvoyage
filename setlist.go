package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Setlist represents data about a setlist including how long it should remain available.
// For temporary setlists, their expiry will be set to creation time plus the
// temporaryPlaylistLifespan. For persisted setlists, their expiry may not be set or may
// be set to a yet to be determined lifespan. (TODO)
type Setlist struct {
	ID     primitive.ObjectID `json:"-"`
	Name   string             `json:"name"`
	Expiry time.Time          `json:"-"`
	Songs  []*Song            `json:"songs"`
}

// Song represents data about a particular song
type Song struct {
	Artist string `json:"artist"`
	Name   string `json:"name"`
}

func setlist(ctx context.Context, db finder, name []byte) (*Setlist, error) {
	return db.find(ctx, string(name))
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
