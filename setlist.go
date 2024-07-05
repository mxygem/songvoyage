package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	tempSetlistName = "temp"
)

// Setlist represents data about a setlist including how long it should remain available.
// For temporary setlists, their expiry will be set to creation time plus the
// temporaryPlaylistLifespan. For persisted setlists, their expiry may not be set or may
// be set to a yet to be determined lifespan. (TODO)
type Setlist struct {
	ID     primitive.ObjectID `json:"-"`
	Name   string             `json:"name"`
	Expiry time.Time          `json:"-"`
	Songs  []*Song            `json:"songs,omitempty"`
}

// Song represents data about a particular song
type Song struct {
	Artist string `json:"artist"`
	Name   string `json:"name"`
}

// TODO: update to include collection
func setlist(ctx context.Context, db finderCreator, name []byte) (*Setlist, error) {
	slName := string(name)
	// if no name is provided, try looking up the temporary setlist
	if slName == "" {
		slName = tempSetlistName
	}

	sl, err := db.find(ctx, slName)
	if err != nil {
		return nil, fmt.Errorf("looking up setlist: %w", err)
	}
	// if found, return the setlist
	if sl != nil {
		return sl, err
	}

	// if the setlist wasn't found and we're looking for the temp setlist, create it.
	// TODO: create & send expiry
	if slName == tempSetlistName {
		sl, err = db.create(ctx, slName)
		if err != nil {
			return nil, fmt.Errorf("creating temp setlist: %w", err)
		}
	}

	return sl, nil
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
