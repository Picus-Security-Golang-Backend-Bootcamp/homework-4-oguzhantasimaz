package infrastructure

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Error definitions
var (
	CannotConnectToDatabaseError = errors.New("Cannot connect to database")
)

// ThrowCannotConnectToDatabaseError throws CannotConnectToDatabaseError error with the given error
func ThrowCannotConnectToDatabaseError(actualErr error) error {
	return errors.New(fmt.Sprintf("%s ==> %s", CannotConnectToDatabaseError.Error(), actualErr.Error()))
}

// NewMySQLDB creates a new MySQLDB instance
func NewMySQLDB(conString string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		PrepareStmt: true, // sonraki sorgular için cache
		NowFunc: func() time.Time {
			return time.Now()
		},
	})

	if err != nil {
		return nil, ThrowCannotConnectToDatabaseError(err)
	}

	return db, nil
}
