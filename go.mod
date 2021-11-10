module stonehenge

go 1.17

// This package is required to create UUID identifiers
require github.com/google/uuid v1.3.0

// These packages are required for the database access
require (
	cloud.google.com/go v0.97.0 // indirect
	cloud.google.com/go/firestore v1.6.0 // indirect
	cloud.google.com/go/storage v1.18.2 // indirect
	firebase.google.com/go v3.13.0+incompatible // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-jwt/jwt/v4 v4.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/net v0.0.0-20210913180222-943fd674d43e // indirect
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1 // indirect
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/api v0.59.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211016002631-37fc39342514 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

// This is required to enable HTTP routing
require github.com/go-chi/chi/v5 v5.0.4 // indirect

require github.com/go-sql-driver/mysql v1.6.0 // indirect
