go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
=> to install gorm drivers for postgres

gorm is used as an ORM, so we are not writing direct SQL queries anymore, we are passing queries through GORM.
gorm, like many other ORMs, pluralizes the struct name to determine the table name. In this case, since the struct being used is Todo{}, GORM will pluralize it to "todos" and use that as the table name. 

Project layout:
```
├── cmd/
│   └── main.go
└── pkg/
    └── todo/
        ├── handler.go
        ├── model.go
        └── repository.go
```