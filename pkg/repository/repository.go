package repository

import (
	segment "github.com/pasha1coil/testingavito/pkg/service/system"

	"github.com/jmoiron/sqlx"
)

type Commands interface {
	CreateUser(user segment.User) (string, error)
	CreateSegment(segment segment.Segment) (string, error)
	DelSegment(segment segment.Segment) (string, error)
	InsertSemUser(NameSegment []string, UserID int) ([]int, error)
	DeleteSemUser(NameSegment []string, UserID int) (string, error)
	GetActiveSlugs(user segment.User) ([]segment.ListNames, error)
	GetSlugHistory(userId int, startDate string, endDate string) ([]segment.History, error)
	Tableinitialization(name segment.DataBase) ([]segment.DbOutput, error)
}

type Repository struct {
	Commands
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Commands: NewAddDb(db),
	}
}
