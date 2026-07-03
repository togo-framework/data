// Package data is togo's data capability — a uniform Query interface over a
// pluggable source. The built-in `pg` backend queries the app's Postgres (the
// togo-framework/db superpowers image: pg_analytics for OLAP + pg_search for BM25
// come for free, since they're just SQL against the same database). Federation
// backends query external sources:
//
//	data-bigquery    Google BigQuery
//	data-databricks  Databricks SQL
//	data-iceberg     Apache Iceberg (open tables on object storage)
//
// Select one with `togo provider:use data <name>` (or TOGO_DATA_PROVIDER). This
// repo is the MAIN plugin (contract + the `pg` default).
package data

import (
	"context"

	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

// Row is one result row (column name → value).
type Row = map[string]any

// Data runs a read query against the active source and returns rows.
type Data interface {
	Query(ctx context.Context, query string, args ...any) ([]Row, error)
}

func init() {
	togo.RegisterProviderFunc("data", togo.PriorityService, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapData, "pg", &pgData{k: k}, true)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "data")
		}
		return nil
	})
}

// FromKernel returns the active Data backend (or nil if none registered).
func FromKernel(k *togo.Kernel) Data {
	if v, ok := k.Get(providers.CapData); ok {
		if d, ok := v.(Data); ok {
			return d
		}
	}
	return nil
}
