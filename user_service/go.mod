module github.com/kristijanpill/go-realworld-example-app/user_service

go 1.18

replace github.com/kristijanpill/go-realworld-example-app/common => ../common

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.2
	github.com/kristijanpill/go-realworld-example-app/common v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
	gorm.io/gorm v1.23.5
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220601144221-27df5f98adab // indirect
	gorm.io/driver/postgres v1.3.7 // indirect
)
