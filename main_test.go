package main

import (
	"testing"

	"github.com/csoban373/testing_examples/db"
	"github.com/csoban373/testing_examples/model"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetPeople(t *testing.T) {
	tests := []struct {
		name     string
		seeds    []model.Person
		expected []model.Person
	}{
		{
			name: "gets all people",
			seeds: []model.Person{
				{ID: 1, Name: "John"},
				{ID: 2, Name: "Jane"},
			},
			expected: []model.Person{
				{ID: 1, Name: "John"},
				{ID: 2, Name: "Jane"},
			},
		},
		{
			name:     "empty DB",
			seeds:    []model.Person{},
			expected: []model.Person{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDb := db.SetupTestDB(t)
			r := require.New(t)
			db.WithinTransaction(t, testDb, func(t *testing.T, tx *gorm.DB) {
				for _, v := range tt.seeds {
					r.NoError(tx.Create(&v).Error)
				}
				result, err := GetPeople(tx)
				r.NoError(err)
				r.Equal(tt.expected, result)
			})
		})
	}
}
