<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/data</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/data"><img src="https://pkg.go.dev/badge/github.com/togo-framework/data.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/data
```

<!-- /togo-header -->

togo's **data** capability — a uniform `Query` interface over a pluggable source.
The built-in **`pg`** backend queries the app's Postgres (the
[db](https://github.com/togo-framework/db) superpowers image), so **`pg_analytics`
(OLAP)** and **`pg_search` (BM25)** come for free — no separate engine. Federation
backends query external sources:

| Backend | Repo |
|---|---|
| pg *(default)* | built in |
| BigQuery | `togo-framework/data-bigquery` |
| Databricks SQL | `togo-framework/data-databricks` |
| Apache Iceberg | `togo-framework/data-iceberg` |

```bash
togo install togo-framework/data-bigquery
togo provider:use data bigquery
```

```go
d := data.FromKernel(k)
rows, _ := d.Query(ctx, "SELECT count(*) AS n FROM orders")
```

> When to reach past `pg`: only when data must out-scale, out-live, or be shared
> beyond a single Postgres — see the [db](https://github.com/togo-framework/db)
> notes. Most apps stay on `pg`.
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
