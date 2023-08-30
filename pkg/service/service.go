package service

import (
	"github.com/pasha1coil/testingavito/pkg/repository"
	segment "github.com/pasha1coil/testingavito/pkg/service/enty"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Commands interface {
	CreateUser(user segment.User) (string, error)
	CreateSegment(segment segment.Segment) (string, error)
	DelSegment(segment segment.Segment) (string, error)
	InsertSemUser(NameSegment []string, UserID int) ([]int, error)
	DeleteSemUser(NameSegment []string, UserID int) (string, error)
	GetActiveSlugs(user segment.User) ([]segment.ListNames, error)
	GetSlugHistory(userId int, startDate string, endDate string) ([]segment.History, error)
	GetCsvHistory(userId int, startDate string, endDate string) (string, error)
	Tableinitialization(name segment.DataBase) ([]segment.DbOutput, error)
}

type Service struct {
	Commands
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Commands: NewAddService(repos.Commands),
	}
}
