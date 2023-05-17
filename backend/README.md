## go migrate

```
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

```
migrate create -ext sql -dir migrations -seq <table name>
```

```
migrate -path migrations -database "mysql://root:root@tcp(db-container:3306)/tm?multiStatements=true" up 2
```

## curl

```
curl -X <METHOD> "http://localhost:8080/task/<method>" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjEsImFkbWluIjpmYWxzZSwiZXhwIjoxNjg0Mjk4NTUzfQ.ejvxdnlYbbwVVuLiEg9wjl7HKOLVybpX2aT7m3v3JAU"
```