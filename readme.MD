go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
=> to install gorm drivers for postgres


Project layout:
```
├── cmd/
│   └── main.go
└── pkg/
    └── auth/
        ├── auth.go
        └── jwt.go
    └── todo/
        ├── handler.go
        ├── model.go
        └── repository.go
    └── user/
        ├── handler.go
        ├── model.go
        ├── tokencontroller.go
        └── usercontroller.go
```

