package data

import (
	"context"
	"database/sql"

	"github.com/togo-framework/togo"
)

// sqlData queries the app's own database through the kernel, whatever its driver
// (SQLite, Postgres, MySQL). On the togo db image (ParadeDB) OLAP (pg_analytics)
// and BM25 search (pg_search) come through the same connection — no separate
// engine. This is the default `data` backend.
type sqlData struct{ k *togo.Kernel }

func (p *sqlData) Query(ctx context.Context, query string, args ...any) ([]Row, error) {
	db, err := p.k.SQL(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

// scanRows reads all rows into []Row (column name → value).
func scanRows(rows *sql.Rows) ([]Row, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var out []Row
	for rows.Next() {
		vals := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}
		row := make(Row, len(cols))
		for i, c := range cols {
			v := vals[i]
			if b, ok := v.([]byte); ok { // normalize driver []byte → string
				v = string(b)
			}
			row[c] = v
		}
		out = append(out, row)
	}
	return out, rows.Err()
}
