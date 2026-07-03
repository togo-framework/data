<!-- togo-brand -->
<p align="center"><img src=".github/assets/togo-mark.svg" width="96" alt="togo" /></p>
<h1 align="center">data</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

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

MIT © fadymondy
