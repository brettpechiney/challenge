$startCockroach = {
    cockroach start --insecure --listen-addr=localhost
}
Start-Job -ScriptBlock $startCockroach

Start-Sleep -s 5

start powershell {Get-Content .\scripts\sql\db.changelog-0.1.0.sql | cockroach sql --insecure --host=localhost:26257}
