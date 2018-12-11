@echo off

start docker-compose up cockroach
TIMEOUT 3 >NUL
cd ./scripts/sql
start sql-migrate up