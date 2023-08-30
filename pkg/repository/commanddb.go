package repository

import (
	"fmt"
	"strings"
	"time"

	segment "github.com/pasha1coil/testingavito/pkg/service/enty"

	"github.com/jmoiron/sqlx"
)

type AddDb struct {
	db *sqlx.DB
}

func NewAddDb(db *sqlx.DB) *AddDb {
	return &AddDb{db: db}
}

func (r *AddDb) CreateUser(user segment.User) (string, error) {
	var message string
	if user.User_number <= 0 {
		return "User number must > 0", nil
	}
	query := fmt.Sprintf("INSERT INTO %s (User_number) values ($1) RETURNING User_number", Users)
	row := r.db.QueryRow(query, user.User_number)
	if err := row.Scan(&message); err != nil {
		return "", err
	}
	return message, nil
}

// func (r *AddDb) CreateSegment(segment segment.Segment) (string, error) {
// 	var name string
// 	query := fmt.Sprintf("INSERT INTO %s (slug_name) values ($1) RETURNING slug_name", slugs)
// 	row := r.db.QueryRow(query, segment.Name)
// 	if err := row.Scan(&name); err != nil {
// 		return "", err
// 	}
// 	return name, nil
// }

func (r *AddDb) CreateSegment(namesegment segment.Segment) (string, error) {
	var name string
	var id = []segment.Segment{}
	query := fmt.Sprintf("INSERT INTO %s (slug_name) values ($1) RETURNING slug_name", slugs)
	row := r.db.QueryRow(query, namesegment.Name)
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	if namesegment.Percent > 0 {
		query := fmt.Sprintf("SELECT user_number FROM %s", Users)
		row, err := r.db.Query(query)
		if err != nil {
			return "", err
		}
		for row.Next() {
			var s segment.Segment
			if err := row.Scan(&s.ID); err != nil {
				return "", err
			}
			id = append(id, s)
		}
		if len(id) > 1 {
			count := float64(len(id)) * float64((namesegment.Percent / 100.0))
			if count <= 1 {
				err := r.InsertSemUser2(namesegment, id[0])
				if err != nil {
					return "", err
				}
			} else {
				for j := 0; j < int(count); j++ {
					err := r.InsertSemUser2(namesegment, id[j])
					if err != nil {
						return "", err
					}
				}
			}
		} else if len(id) == 0 {
			r.DelSegment(namesegment)
			return "Add User,pls", nil
		} else {
			err := r.InsertSemUser2(namesegment, id[0])
			if err != nil {
				return "", err
			}
		}
	} else if namesegment.Percent < 0 {
		r.DelSegment(namesegment)
		return "Fix Percent ,pls", nil
	}
	return name, nil
}

// res := r.AddInHistory2(segment)
// 	if res != nil {
// 		return false, res
// 	}
// 	query1 := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1)", UsersSlug)
// 	_, err := r.db.Exec(query1, segment.Name)
// 	if err != nil {
// 		return false, err
// 	}
// 	query2 := fmt.Sprintf("DELETE FROM %s WHERE slug_name=($1)", slugs)
// 	_, err = r.db.Exec(query2, segment.Name)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil

func (r *AddDb) DelSegment(slug segment.Segment) (string, error) {
	var array = []segment.ListNames{}
	rows, err := r.db.Query("SELECT slug_name FROM slugs WHERE slug_name=($1)", slug.Name)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		var s segment.ListNames
		if err := rows.Scan(&s.Name); err != nil {
			return "", err
		}
		array = append(array, s)
	}
	if len(array) > 0 {
		res := r.AddInHistory2(slug)
		if res != nil {
			return "", res
		}
		query1 := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1)", UsersSlug)
		_, err := r.db.Exec(query1, slug.Name)
		if err != nil {
			return "", err
		}
		query2 := fmt.Sprintf("DELETE FROM %s WHERE slug_name=($1)", slugs)
		_, err = r.db.Exec(query2, slug.Name)
		if err != nil {
			return "", err
		}
	} else {
		return "No such data", nil
	}
	return "OK", nil
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

func (r *AddDb) InsertSemUser2(NameSegment segment.Segment, UserID segment.Segment) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1) AND UserID = ($2)", UsersSlug)
	_, err := r.db.Exec(query, NameSegment.Name, UserID.ID)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("INSERT INTO %s (name_slug,UserID) values ($1,$2) RETURNING id", UsersSlug)
	_, err = r.db.Exec(query, NameSegment.Name, UserID.ID)
	if err != nil {
		return err
	}
	if r.AddInHistory(NameSegment.Name, UserID.ID, "ADD"); err != nil {
		return err
	}
	return nil
}

//DeleteSemUser

func (r *AddDb) DeleteSemUser(NameSegment []string, UserID int) (string, error) {
	var id = []segment.TableSlugs{}
	if len(NameSegment) > 0 {
		for _, i := range NameSegment {

			rows, err := r.db.Query("SELECT id FROM UsersSlug WHERE name_slug=($1) AND UserID= ($2)", i, UserID)
			if err != nil {
				return "", err
			}
			defer rows.Close()
			for rows.Next() {
				var s segment.TableSlugs
				if err := rows.Scan(&s.ID); err != nil {
					return "", err
				}
				id = append(id, s)
			}
		}
		if len(id) > 0 {
			for _, i := range NameSegment {

				r.AddInHistory(i, UserID, "DELETE")
				query := fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1) AND UserID= ($2)", UsersSlug)
				_, err := r.db.Exec(query, i, UserID)
				if err != nil {
					return "", err
				}
			}
		} else {
			return "No such data", nil
		}
	} else {
		return "Empty list value", nil
	}
	return "OK", nil
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

func (r *AddDb) AddInHistory2(NameSegment segment.Segment) error {
	mode := "DELETE"
	var id = []segment.Segment{}
	query := fmt.Sprintf("SELECT userid FROM %s WHERE name_slug=($1)", History)
	row, err := r.db.Query(query, NameSegment.Name)
	if err != nil {
		return err
	}
	for row.Next() {
		var s segment.Segment
		if err := row.Scan(&s.ID); err != nil {
			return err
		}
		id = append(id, s)
	}
	for _, i := range id {
		query = fmt.Sprintf("DELETE FROM %s WHERE name_slug=($1) AND UserID = ($2) AND mode = ($3)", History)
		_, err = r.db.Exec(query, NameSegment.Name, i.ID, mode)
		if err != nil {
			return err
		}

		query = fmt.Sprintf("INSERT INTO %s (name_slug,UserID,mode,created) values ($1,$2,$3,$4)", History)
		_, err = r.db.Exec(query, NameSegment.Name, i.ID, mode, time.Now())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *AddDb) Tableinitialization(name segment.DataBase) ([]segment.DbOutput, error) {
	var array = []segment.DbOutput{}
	name.Name = strings.ToLower(name.Name)
	if name.Name == "users" {
		rows, err := r.db.Query("SELECT * FROM Users")
		if err != nil {
			return array, err
		}
		defer rows.Close()
		for rows.Next() {
			var s segment.DbOutput
			if err := rows.Scan(&s.User_id); err != nil {
				return array, err
			}
			array = append(array, s)
		}
	} else if name.Name == "slugs" {
		rows, err := r.db.Query("SELECT * FROM slugs")
		if err != nil {
			return array, err
		}
		defer rows.Close()
		for rows.Next() {
			var s segment.DbOutput
			if err := rows.Scan(&s.Name); err != nil {
				return array, err
			}
			array = append(array, s)
		}
	} else if name.Name == "usersslug" {
		rows, err := r.db.Query("SELECT * FROM UsersSlug")
		if err != nil {
			return array, err
		}
		defer rows.Close()
		for rows.Next() {
			var s segment.DbOutput
			if err := rows.Scan(&s.ID, &s.User_id, &s.Name); err != nil {
				return array, err
			}
			array = append(array, s)
		}
	} else if name.Name == "history" {
		rows, err := r.db.Query("SELECT * FROM History")
		if err != nil {
			return array, err
		}
		defer rows.Close()
		for rows.Next() {
			var s segment.DbOutput
			if err := rows.Scan(&s.ID, &s.User_id, &s.Name, &s.Mode, &s.Created); err != nil {
				return array, err
			}
			array = append(array, s)
		}
	} else {
		return array, nil
	}
	return array, nil
}
