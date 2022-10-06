https://github.com/go-redis/redis

go-redis supports 2 last Go versions and requires a Go version with modules support. So make sure to initialize a Go module:

```bash
go mod init github.com/my/repo
```

If you are using Redis 6, install go-redis/v8:
```bash
go get github.com/go-redis/redis/v8
```

If you are using Redis 7, install go-redis/v9:
```bash
go get github.com/go-redis/redis/v9
```
