package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSetlist(t *testing.T) {
	testTime := time.Now()
	testID := primitive.NewObjectIDFromTimestamp(testTime)

	testCases := []struct {
		name        string
		setlistName string
		db          func(t *testing.T) *Mockdber
		expected    *Setlist
		expectedErr error
	}{
		{
			name:        "no name provided - searches for and returns the temp setlist when exists",
			setlistName: "",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, tempSetlistName).Return(
					&Setlist{
						Name:   tempSetlistName,
						Expiry: testTime.Add(24 * time.Hour),
					},
					nil,
				)

				return db
			},
			expected: &Setlist{
				Name:   "temp",
				Expiry: testTime.Add(24 * time.Hour),
			},
		},
		{
			name: "no name provided - searches for temp setlist and creates if not exists",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, tempSetlistName).Return(
					nil,
					nil,
				)
				db.On("create", mock.Anything, tempSetlistName).Return(
					&Setlist{
						Name:   tempSetlistName,
						Expiry: testTime.Add(24 * time.Hour),
					},
					nil,
				)

				return db
			},
			expected: &Setlist{
				Name:   tempSetlistName,
				Expiry: testTime.Add(24 * time.Hour),
			},
		},
		{
			name:        "name provided - returns matched playlist",
			setlistName: "Doomed Fingers",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, "Doomed Fingers").Return(
					&Setlist{
						ID:   testID,
						Name: "Doomed Fingers",
						Songs: []*Song{
							{Artist: "Dragonforce", Name: "Through the Fire and Flames"},
						},
					},
					nil,
				)

				return db
			},
			expected: &Setlist{
				ID:   testID,
				Name: "Doomed Fingers",
				Songs: []*Song{
					{Artist: "Dragonforce", Name: "Through the Fire and Flames"},
				},
			},
		},
		{
			name:        "name provided - returns double nil when unmatched",
			setlistName: "Djent Madness",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, "Djent Madness").Return(
					nil,
					nil,
				)

				return db
			},
			expected:    nil,
			expectedErr: nil,
		},
		{
			name:        "name provided - db returns error when looking up setlist",
			setlistName: "Doomed Fingers",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, "Doomed Fingers").Return(
					nil,
					fmt.Errorf("it broke"),
				)

				return db
			},
			expectedErr: fmt.Errorf("looking up setlist: it broke"),
		},
		{
			name: "name not provided - db returns error when creating temp setlist",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, tempSetlistName).Return(
					nil,
					nil,
				)
				db.On("create", mock.Anything, tempSetlistName).Return(
					nil,
					fmt.Errorf("setlist too epic"),
				)

				return db
			},
			expectedErr: fmt.Errorf("creating temp setlist: setlist too epic"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			db := tc.db(t)

			actual, err := setlist(ctx, db, []byte(tc.setlistName))

			assert.Equal(t, tc.expected, actual)
			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
