package data

import (
	"database/sql"
	"testing"

	"github.com/togo-framework/togo"
	_ "modernc.org/sqlite"
)

func TestScanRows(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if _, err := db.Exec(`CREATE TABLE t(id INTEGER, name TEXT); INSERT INTO t VALUES (1,'a'),(2,'b')`); err != nil {
		t.Fatal(err)
	}
	rows, err := db.Query("SELECT id, name FROM t ORDER BY id")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	got, err := scanRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0]["name"] != "a" || got[1]["name"] != "b" {
		t.Fatalf("scanRows = %v", got)
	}
}

func TestFromKernelDefaultsToPg(t *testing.T) {
	d := FromKernel(togo.New())
	if d == nil {
		t.Fatal("no data backend registered")
	}
	if _, ok := d.(*pgData); !ok {
		t.Fatalf("default should be pgData, got %T", d)
	}
}
