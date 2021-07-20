package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func New(connString string) (*DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func ConnString(host string, port int, user, password, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
}

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (d *DB) GetUserByName(name string) ([]*User, error) {
	stmt, err := d.Prepare("select * from users where name=$1")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}
	users := []*User{}
	for rows.Next() {
		r := User{}
		if err = rows.Scan(&r.ID, &r.Name, &r.Age, &r.Profession, &r.Friendly); err != nil {
			return nil, err
		}
		users = append(users, &r)
	}

	return users, nil
}
