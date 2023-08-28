package service

import (
	segment "github.com/pasha1coil/testingavito"
	"github.com/pasha1coil/testingavito/pkg/repository"
)

type Commands interface {
	CreateUser(user segment.User) (int, error)
	CreateSegment(segment segment.Segment) (string, error)
	DelSegment(segment segment.Segment) (bool, error)
	InsertSemUser(NameSegment []string, UserID int) ([]int, error)
	DeleteSemUser(NameSegment []string, UserID int) (bool, error)
	GetActiveSlugs(user segment.User) ([]segment.ListNames, error)
	GetSlugHistory(userId int, startDate string, endDate string) ([]segment.History, error)
	GetCsvHistory(userId int, startDate string, endDate string) (string, error)
}

type Service struct {
	Commands
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Commands: NewAddService(repos.Commands),
	}
}
