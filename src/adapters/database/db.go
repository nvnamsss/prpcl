package database

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBAdapter interface {
	Open(connectionString string, config gorm.Config) error
	Begin() DBAdapter
	RollbackUselessCommitted()
	Commit()
	Close()
	Gormer() *gorm.DB
	DB() (*sql.DB, error)
}

type adapter struct {
	gormer      *gorm.DB
	isCommitted bool
}

// NewDB returns a new instance of DB.
func NewDatabase() DBAdapter {
	return &adapter{}
}

// Open opens a DB connection.
func (db *adapter) Open(connectionString string, config gorm.Config) error {
	gormDB, err := gorm.Open(mysql.Open(connectionString), &config)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	db.gormer = gormDB
	return nil
}

// Begin starts a DB transaction.
func (db *adapter) Begin() DBAdapter {
	tx := db.gormer.Begin()
	return &adapter{
		gormer:      tx,
		isCommitted: false,
	}
}

// RollbackUselessCommitted rollbacks useless DB transaction committed.
func (db *adapter) RollbackUselessCommitted() {
	if !db.isCommitted {
		db.gormer.Rollback()
	}
}

// Commit commits a DB transaction.
func (db *adapter) Commit() {
	if !db.isCommitted {
		db.gormer.Commit()
		db.isCommitted = true
	}
}

// Close closes DB connection.
func (db *adapter) Close() {
	sqlDB, err := db.gormer.DB()
	if err != nil {
		return
	}

	_ = sqlDB.Close()
}

// Gormer returns an instance of gorm.DB.
func (db *adapter) Gormer() *gorm.DB {
	return db.gormer
}

// DB returns an instance of sql.DB.
func (db *adapter) DB() (*sql.DB, error) {
	return db.gormer.DB()
}
