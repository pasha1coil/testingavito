package repository

import (
	"fmt"
	segment "testingavito"
	"time"

	"github.com/jmoiron/sqlx"
)

type AddDb struct {
	db *sqlx.DB
}

func NewAddDb(db *sqlx.DB) *AddDb {
	return &AddDb{db: db}
}

func (r *AddDb) CreateUser(user segment.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (User_number) values ($1) RETURNING User_number", Users)
	row := r.db.QueryRow(query, user.User_number)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AddDb) CreateSegment(segment segment.Segment) (string, error) {
	var name string
	query := fmt.Sprintf("INSERT INTO %s (slug_name) values ($1) RETURNING slug_name", slugs)
	row := r.db.QueryRow(query, segment.Name)
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}

func (r *AddDb) DelSegment(segment segment.Segment) (bool, error) {
	query1 := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1)", UsersSlug)
	_, err := r.db.Exec(query1, segment.Name)
	if err != nil {
		return false, err
	}
	query2 := fmt.Sprintf("DELETE FROM %s WHERE slug_name=($1)", slugs)
	_, err = r.db.Exec(query2, segment.Name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AddDb) InsertSemUser(NameSegment []string, UserID int) ([]int, error) {
	var id []int
	var ids int
	for _, i := range NameSegment {
		query := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1) AND UserID = ($2)", UsersSlug)
		_, err := r.db.Exec(query, i, UserID)
		if err != nil {
			return id, err
		}
		query = fmt.Sprintf("INSERT INTO %s (name_slug,UserID) values ($1,$2) RETURNING id", UsersSlug)
		rows := r.db.QueryRow(query, i, UserID)
		if err := rows.Scan(&ids); err != nil {
			return id, err
		}
		id = append(id, ids)
		if r.AddInHistory(i, UserID, "ADD"); err != nil {
			return id, err
		}
	}
	return id, nil
}

//DeleteSemUser

func (r *AddDb) DeleteSemUser(NameSegment []string, UserID int) (bool, error) {
	var id int
	for _, i := range NameSegment {

		query := fmt.Sprintf("SELECT id FROM %s WHERE name_slug=($1) AND UserID= ($2)", UsersSlug)
		row := r.db.QueryRow(query, i, UserID)
		if err := row.Scan(&id); err != nil {
			continue
		}
		if id != 0 {
			r.AddInHistory(i, UserID, "DELETE")
			query := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1) AND UserID= ($2)", UsersSlug)
			_, err := r.db.Exec(query, i, UserID)
			if err != nil {
				return false, err
			}
		} else {
			return false, nil
		}
	}
	return true, nil
}

// GetActiveSlugs
func (r *AddDb) GetActiveSlugs(user segment.User) ([]segment.ListNames, error) {
	var names = []segment.ListNames{}
	rows, err := r.db.Query("SELECT name_slug FROM UsersSlug WHERE UserID =($1)", user.User_number)
	if err != nil {
		return names, err
	}
	defer rows.Close()
	for rows.Next() {
		var s segment.ListNames
		if err := rows.Scan(&s.Name); err != nil {
			return names, err
		}
		names = append(names, s)
	}
	return names, err
}

func (r *AddDb) AddInHistory(NameSegment string, UserID int, mode string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1) AND UserID = ($2) AND mode = ($3)", History)
	_, err := r.db.Exec(query, NameSegment, UserID, mode)
	if err != nil {
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (name_slug,UserID,mode,created) values ($1,$2,$3,$4)", History)
	_, err = r.db.Exec(query, NameSegment, UserID, mode, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *AddDb) GetSlugHistory(userId int, startDate string, endDate string) ([]segment.History, error) {
	layout := "2006/01/02"
	s, _ := time.Parse(layout, startDate+"/02")
	NewstartDate := s.Format("2006-01-02")
	d, _ := time.Parse(layout, endDate+"/02")
	NewendDate := d.Format("2006-01-02")
	fmt.Println(NewstartDate, NewendDate)
	rows, err := r.db.Query("SELECT UserID, name_slug, mode, created FROM History WHERE UserID=($1) AND created >= ($2) AND created <= ($3)",
		userId, NewstartDate, NewendDate)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var slugHistory = []segment.History{}
	for rows.Next() {
		var s segment.History
		if err := rows.Scan(&s.User_id, &s.Name, &s.Mode, &s.Time); err != nil {
			return slugHistory, err
		}
		slugHistory = append(slugHistory, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return slugHistory, nil
}
