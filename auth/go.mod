module local/auth

go 1.16

require (
	github.com/go-redis/redis/v8 v8.8.2
	github.com/google/go-cmp v0.5.5
	github.com/gorilla/mux v1.8.0
	github.com/markbates/goth v1.67.1
	local/common v1.0.0
)

replace local/common => ../common
