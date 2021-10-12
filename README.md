# Content Structure

```
.
├── base
│   ├── api
│   │   ├── base.proto
│   │   ├── base.sh
│   │   └── gen
│   │       └── go
│   │           ├── base_grpc.pb.go
│   │           └── base.pb.go
│   ├── base
│   │   └── base.go
│   ├── conf
│   │   └── envs
│   │       ├── db.env
│   │       └── test.env
│   └── main.go
├── errs
│   └── logerrs.go
├── gateway
│   └── main.go
├── go.mod
├── go.sum
├── LICENSE
├── models
│   └── paste.go
├── paste
│   ├── api
│   │   ├── gen
│   │   │   └── go
│   │   │       ├── paste_grpc.pb.go
│   │   │       ├── paste.pb.go
│   │   │       └── paste.pb.gw.go
│   │   ├── paste.proto
│   │   ├── paste.sh
│   │   └── paste.yaml
│   ├── main.go
│   └── paste
│       ├── paste.go
│       └── transfer.go
├── README.md
├── repositories
│   ├── paste_repo.go
│   └── user_repo.go
├── server
│   └── grpc.go
├── tests
│   └── random_test.go
└── user
    ├── api
    │   ├── gen
    │   │   └── go
    │   │       ├── user_grpc.pb.go
    │   │       ├── user.pb.go
    │   │       └── user.pb.gw.go
    │   ├── user.proto
    │   ├── user.sh
    │   └── user.yaml
    ├── main.go
    └── user
        ├── transfer.go
        └── user.go
```



## TODO

- login achieve with jwt
- password encrype
- add Redis

