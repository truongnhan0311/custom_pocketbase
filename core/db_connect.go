//go:build !no_default_driver

package core

import (
	"database/sql"
	"github.com/pocketbase/dbx"
	"modernc.org/sqlite"
)

func init() {
	sql.Register("sqlite3", &sqlite.Driver{})
}

func DefaultDBConnect(dbPath string) (*dbx.DB, error) {
	// Note: the busy_timeout pragma must be first because
	// the connection needs to be set to block on busy before WAL mode
	// is set in case it hasn't been already set by another connection.
	pragmas := "?_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=journal_size_limit(200000000)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(ON)&_pragma=temp_store(MEMORY)&_pragma=cache_size(-16000)"

	db, err := dbx.Open("sqlite3", dbPath+pragmas)
	if err != nil {
		return nil, err
	}

	return db, nil
}
