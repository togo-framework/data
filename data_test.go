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

func TestFromKernelDefaultsToSQL(t *testing.T) {
	d := FromKernel(togo.New())
	if d == nil {
		t.Fatal("no data backend registered")
	}
	if _, ok := d.(*sqlData); !ok {
		t.Fatalf("default should be sqlData, got %T", d)
	}
}

// The default backend queries the app's database on any driver, so it is not
// named after Postgres (it logged backend=pg on SQLite); "pg" still selects it.
func TestDefaultBackendName(t *testing.T) {
	for active, want := range map[string]string{"": "sql", "sql": "sql", "pg": "pg", "bigquery": "sql"} {
		if got := defaultBackend(active); got != want {
			t.Errorf("defaultBackend(%q) = %q, want %q", active, got, want)
		}
	}
}
