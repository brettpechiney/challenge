#!/usr/bin/env bash
cockroach start --insecure --listen-addr=localhost &
cockroach sql --insecure --host=localhost:26257 < ./scripts/sql/db.changelog-0.1.0.sql