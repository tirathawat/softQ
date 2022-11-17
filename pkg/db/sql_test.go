package db_test

import (
	"database/sql"
	"testing"

	"github.com/tirathawat/softQ/pkg/db"
)

type mockDB struct {
	query        string
	lastInsearId int64
	rowAffected  int64
}

func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsearId, nil
}

func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowAffected, nil
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}

func TestExecQuery(t *testing.T) {
	var want int64 = 32
	query := "SELECT * FROM users"
	mock := &mockDB{rowAffected: want}

	r, err := db.ExecQuery(mock, query)
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if r != want {
		t.Errorf("got %v, want %v", r, want)
	}

	if mock.query != query {
		t.Errorf("got %v, want %v", mock.query, query)
	}
}
