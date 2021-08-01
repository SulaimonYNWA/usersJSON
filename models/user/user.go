package user

import (
	"database/sql"
	"errors"
)

type Model interface {
	AddUser(user User) (id int64, err error)
	GetAll() ([]User, error)
	Get(int64) (*User, error)
}

type model struct {
	opts Options
}

func NewModel(opts Options) Model {
	return model{opts: opts}
}

type Options struct {
	DB *sql.DB
}

// crud methods -------------------------------------------

func (m model) AddUser(user User) (id int64, err error) {
	q := "INSERT INTO users (name, surname, age, sex, email, password, isAdmin) VALUES (?, ?, ?, ?, ?, ?, ?)"
	res, err := m.opts.DB.Exec(q, user.Name, user.Surname, user.Age, user.Sex, user.Email, user.Password, user.IsAdmin)
	if err != nil {
		err = errors.New("create user sql insert: " + err.Error())
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		err = errors.New("user last insert id: " + err.Error())
		return
	}

	return
}

func (m model) GetAll() ([]User, error) {
	users := []User{}

	q := "SELECT id, name, surname, age, sex, email, password, isAdmin FROM users"

	rows, err := m.opts.DB.Query(q)
	if err != nil {
		return nil, errors.New("query rows: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		u := &SQLUser{}
		err = rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Age, &u.Sex, &u.Email, &u.Password, &u.IsAdmin)
		if err != nil {
			return nil, errors.New("rows scan: " + err.Error())
		}

		res, _ := sqlToEntity(*u)
		users = append(users, *res)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.New("rows error: " + err.Error())
	}

	return users, nil
}

func (m model) Get(id int64) (*User, error) {

	q := "SELECT id, name, surname, age, sex, email, password, isAdmin FROM users WHERE id = ?"

	u := SQLUser{}
	err := m.opts.DB.QueryRow(q, id).
		Scan(&u.ID, &u.Name, &u.Surname, &u.Age, &u.Sex, &u.Email, &u.Password, &u.IsAdmin)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, errors.New("row scan: ")
	}

	user, _ := sqlToEntity(u)
	return user, nil
}
