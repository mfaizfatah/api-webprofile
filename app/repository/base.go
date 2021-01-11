package repository

import (
	"gorm.io/gorm"
)

// repo struct with value mysqldb connection
type repo struct {
	db *gorm.DB
}

// Repo represent the Repository contract
type Repo interface {
	// insert
	Insert(table string, i interface{}) error
}

/*NewRepo will create an object that represent the Repository interface (Repo)
 * @parameter
 * db - mysql database connection
 *
 * @represent
 * interface Repo
 *
 * @return
 * repo struct with value db (mysql database connection)
 */
func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
