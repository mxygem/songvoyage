package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Create composite interfaces to allow for finding & XYZ or determine alternative
// solution. Maybe middleware that attempts to look up a setlist when handling request?
// finder provides the method find, to be used to look up a setlist by name.
type finder interface {
	find(ctx context.Context, name string) (*Setlist, error)
}

// creator provides the method create, used to create a new setlist by name.
type creator interface {
	create(ctx context.Context, name string) (*Setlist, error)
}

// deleter provides the method delete, used to delete a setlist by name.
type deleter interface {
	delete(ctx context.Context, name string) error
}

// clearer provides the method clear, used to clear a setlist by name.
type clearer interface {
	clear(ctx context.Context, name string) error
}

// saver provides the method save, used to save a temporary setlist as a persisted one
// with the provided name.
type saver interface {
	save(ctx context.Context, name string) (*Setlist, error)
}

// updater provides the method update, used to update a setlist's name from the provided
// old name to the provided new name.
type updater interface {
	update(ctx context.Context, oldName, newName string) (*Setlist, error)
}

// songer provides the methods add & remove, which are used to modify a setlist's list of
// songs.
type songer interface {
	// add updates a setlist's song list, appending the provided artist & song.
	add(ctx context.Context, setlistName, artistName, songName string) error
	// remove updates a setlist's song list, removing a song matching the provided name
	// or position in the song list.
	remove(ctx context.Context, setlistName, songName string, songNumber int) error
}

type dber interface {
	finder
	creator
	deleter
	clearer
	saver
	updater
	songer
}

// db represents the accesor to the server's database and implements the core interfaces
// above.
type db struct {
	client *mongo.Client
}

// newDB returns a new instance of db with a configured client.
func newDB() *db {
	return &db{
		client: nil, // TODO: setup mongo connection
	}
}

func (db *db) find(ctx context.Context, name string) (*Setlist, error) {
	return nil, nil
}

func (db *db) create(ctx context.Context, name string) (*Setlist, error) {
	return nil, nil
}

func (db *db) delete(ctx context.Context, name string) error {
	return nil
}

func (db *db) clear(ctx context.Context, name string) error {
	return nil
}

func (db *db) save(ctx context.Context, name string) (*Setlist, error) {
	return nil, nil
}

func (db *db) update(ctx context.Context, oldName, newName string) (*Setlist, error) {
	return nil, nil
}

func (db *db) add(ctx context.Context, setlistName, artistName, songName string) error {
	return nil
}

func (db *db) remove(ctx context.Context, setlistName, songName string, songNumber int) error {
	return nil
}
