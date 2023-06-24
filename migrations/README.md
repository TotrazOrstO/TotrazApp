# МИГРАЦИИ

### Накатить новую миграцию

 `migrate -database 'postgres://totalim_chat:test@:1234/chat?sslmode=disable' -path ./db/migrations up 1`