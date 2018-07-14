package infrastructure

import (
	"blogo-server/api/interfaces/database"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler - sql connection handler
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler - init sql handler
func NewSQLHandler() *SQLHandler {
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	dbconf := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

// Execute - exec
func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

// Query - query
func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

// SQLResult - sql result
type SQLResult struct {
	Result sql.Result
}

// LastInsertId - insert id
func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected - aff
func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// SQLRow - sql
type SQLRow struct {
	Rows *sql.Rows
}

// Scan - scan
func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next - next
func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

// Close - close
func (r SqlRow) Close() error {
	return r.Rows.Close()
}
