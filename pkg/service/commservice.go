package service

import (
	"fmt"
	"os"
	"time"

	"github.com/pasha1coil/testingavito/pkg/repository"
	segment "github.com/pasha1coil/testingavito/pkg/service/enty"
)

type AddService struct {
	repo repository.Commands
}

func NewAddService(repo repository.Commands) *AddService {
	return &AddService{repo: repo}
}

func (a *AddService) CreateUser(user segment.User) (string, error) {
	return a.repo.CreateUser(user)
}

func (a *AddService) CreateSegment(segment segment.Segment) (string, error) {
	return a.repo.CreateSegment(segment)
}

func (a *AddService) DelSegment(segment segment.Segment) (string, error) {
	return a.repo.DelSegment(segment)
}

func (a *AddService) InsertSemUser(NameSegment []string, UserID int) ([]int, error) {
	return a.repo.InsertSemUser(NameSegment, UserID)
}

func (a *AddService) DeleteSemUser(NameSegment []string, UserID int) (string, error) {
	return a.repo.DeleteSemUser(NameSegment, UserID)
}

func (a *AddService) GetActiveSlugs(user segment.User) ([]segment.ListNames, error) {
	return a.repo.GetActiveSlugs(user)
}

func (a *AddService) GetSlugHistory(userId int, startDate string, endDate string) ([]segment.History, error) {
	return a.repo.GetSlugHistory(userId, startDate, endDate)
}

const head = "NAME | MODE | DATE\n"

func (a *AddService) GetCsvHistory(userId int, startDate string, endDate string) (string, error) {
	history, err := a.repo.GetSlugHistory(userId, startDate, endDate)
	if err != nil {
		return "", err
	}
	file := "static/otchet.csv"
	text, err := os.Create("static/otchet.csv")
	if err != nil {
		return "", err
	}
	defer text.Close()
	text.WriteString(head)
	for _, data := range history {
		text.WriteString(fmt.Sprintf(
			"%s,%s,%s\n",
			data.Name,
			data.Mode,
			data.Time.Format(time.DateOnly),
		))
	}
	url := "http://localhost:8080/" + file
	return url, nil
}

//Tableinitialization(name segment.DataBase) ([]segment.DbOutput, error)

func (a *AddService) Tableinitialization(name segment.DataBase) ([]segment.DbOutput, error) {
	return a.repo.Tableinitialization(name)
}
