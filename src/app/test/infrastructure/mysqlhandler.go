package infrastructure

import (
	"database/sql"
	"interfaces/repository/rdb"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() rdb.SqlHandler {
	dbPath := "root:root@tcp(127.0.0.1:3306)/test_db?parseTime=true"
	conn, err := sql.Open("mysql", dbPath)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
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

func DeleteAllRecords(sqlHandler rdb.SqlHandler) error {
	sqlHandler.Execute("SET FOREIGN_KEY_CHECKS = 0;")
	tables := []string{"aliases", "memberships", "contents", "performances", "participations", "external_services", "external_ids", "occupations", "artists", "songs", "albums"}
	for _, table := range tables {
		_, err := sqlHandler.Execute("TRUNCATE TABLE " + table + ";")
		if err != nil {
			return err
		}
	}
	sqlHandler.Execute("SET FOREIGN_KEY_CHECKS = 1;")
	return nil
}
