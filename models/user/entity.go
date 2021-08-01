package user

import "database/sql"

type User struct {
	ID       int64
	Email    string
	Name     string
	Surname  string
	Age      int64
	Sex      string
	Password string
	IsAdmin  bool
}

type SQLUser struct {
	ID       sql.NullInt64
	Email    sql.NullString
	Name     sql.NullString
	Surname  sql.NullString
	Age      sql.NullInt64
	Sex      sql.NullString
	Password sql.NullString
	IsAdmin  sql.NullBool
}

// converter functions ------------------------------------

func sqlToEntity(v SQLUser) (*User, error) {
	u := &User{
		ID:       v.ID.Int64,
		Name:     v.Name.String,
		Surname:  v.Name.String,
		Age:      v.Age.Int64,
		Sex:      v.Sex.String,
		Email:    v.Email.String,
		Password: v.Password.String,
		IsAdmin:  v.IsAdmin.Bool,
	}
	return u, nil
}

func entityToSQL(v User) (*SQLUser, error) {
	u := &SQLUser{
		ID:       sql.NullInt64{Int64: v.ID, Valid: true},
		Name:     sql.NullString{String: v.Name, Valid: true},
		Surname:  sql.NullString{String: v.Surname, Valid: true},
		Age:      sql.NullInt64{Int64: v.Age, Valid: true},
		Sex:      sql.NullString{String: v.Sex, Valid: true},
		Email:    sql.NullString{String: v.Email, Valid: true},
		Password: sql.NullString{String: v.Password, Valid: true},
		IsAdmin:  sql.NullBool{Bool: v.IsAdmin, Valid: true},
	}
	return u, nil
}
