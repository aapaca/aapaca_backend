package infrastructure

import (
	"database/sql"
	"interfaces/repository/rdb"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() rdb.SqlHandler {
	var conn *sql.DB
	if app_env := os.Getenv("APP_ENV"); app_env == "production" {
		conn = NewCloudSqlConnection()
	} else {
		conn = NewMysqlConnection()
	}
	err := conn.Ping()
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func NewMysqlConnection() *sql.DB {
	c := NewMysqlConfig()
	dbPath := c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database + "?parseTime=true"
	conn, err := sql.Open("mysql", dbPath)
	if err != nil {
		panic(err)
	}
	return conn
}

func NewCloudSqlConnection() *sql.DB {
	c := NewCloudSqlConfig()
	dbPath := c.User + ":" + c.Password + "@unix(/cloudsql/" + c.ConnectionName + ")/" + c.Database + "?parseTime=true"
	conn, err := sql.Open("mysql", dbPath)
	if err != nil {
		panic(err)
	}
	return conn
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (rdb.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (rdb.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
