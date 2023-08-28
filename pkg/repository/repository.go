package repository

import (
	segment "github.com/pasha1coil/testingavito"

	"github.com/jmoiron/sqlx"
)

type Commands interface {
	CreateUser(user segment.User) (int, error)
	CreateSegment(segment segment.Segment) (string, error)
	DelSegment(segment segment.Segment) (bool, error)
	InsertSemUser(NameSegment []string, UserID int) ([]int, error)
	DeleteSemUser(NameSegment []string, UserID int) (bool, error)
	GetActiveSlugs(user segment.User) ([]segment.ListNames, error)
	GetSlugHistory(userId int, startDate string, endDate string) ([]segment.History, error)
}

type Repository struct {
	Commands
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Commands: NewAddDb(db),
	}
}
