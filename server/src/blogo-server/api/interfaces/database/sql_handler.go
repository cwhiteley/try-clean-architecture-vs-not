package database

// SQLHandler - sql handler
type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

// Result - result
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Row - row
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
