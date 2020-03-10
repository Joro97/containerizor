package postgres

import (
	"database/sql"
	"fmt"

	"github.com/asankov/containerizor/pkg/models"

	// to register PostreSQL driver
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func (db *Database) CreateUser(user *models.User) error {
	// TODO: do this in one transaction
	// BIGGER TODO: fix sql injection
	sql := fmt.Sprintf("INSERT INTO USERS(username, passwordHash) VALUES ('%s', '%s');", user.Username, user.HashedPassword)
	if _, err := db.db.Exec(sql); err != nil {
		return err
	}
	sql = fmt.Sprintf("CREATE SCHEMA %s;", user.Username)
	if _, err := db.db.Exec(sql); err != nil {
		return err
	}
	// TODO: this could be template
	sql = fmt.Sprintf(`CREATE TABLE %s.CONTAINERS (id TEXT)`, user.Username)
	if _, err := db.db.Exec(sql); err != nil {
		return err
	}
	return nil
}
func (db *Database) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}
func (db *Database) GetUserByUsername(username string) (*models.User, error) {
	return nil, nil
}
func (db *Database) Close() {
	db.Close()
}

// New connects to a PostgreSQL instance with the
// given parameters and returns the connection,
// or an error if such occured.
func New(host string, port int, user string, dbName string) (*Database, error) {
	// TODO: password, sslmode
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, err
}