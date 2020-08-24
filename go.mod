module golang_api

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-delve/delve v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	golang.org/x/tools v0.0.0-20200823205832-c024452afbcd // indirect
	google.golang.org/appengine v1.6.6 // indirect
	local.packages/db v0.0.0-00010101000000-000000000000
	local.packages/tables v0.0.0-00010101000000-000000000000
)

replace local.packages/db => ./db

replace local.packages/tables => ./tables
