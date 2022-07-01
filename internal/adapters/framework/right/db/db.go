package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v\n", err)
	}

	//test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v\n", err)
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v\n", err)
	}
}

func (da Adapter) AddToHistory(ans int32, operation string) error {
	queryString, args, err := sq.Inster("arith_history").Columns("data", "answer", "operation").
		Values(time.Now(), ans, operation).ToSql()
	if err != nil {
		return err
	}
	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
