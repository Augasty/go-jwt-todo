A crud backend with jwt authentication made using go, gin, gorm, go-jwt and postgres.

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

