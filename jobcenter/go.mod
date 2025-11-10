module jobcenter

go 1.25.3

require (
<<<<<<< HEAD
	github.com/go-co-op/gocron v1.37.0
	github.com/segmentio/kafka-go v0.4.49
	github.com/zeromicro/go-zero v1.9.2
	go.mongodb.org/mongo-driver v1.17.6
)

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
=======
	github.com/zeromicro/go-zero v1.7.6
	mscoin-common v0.0.0
)

replace mscoin-common => ../mscoin-common

require (
	github.com/go-co-op/gocron v1.37.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/segmentio/kafka-go v0.4.49 // indirect
>>>>>>> origin/main
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
<<<<<<< HEAD
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/automaxprocs v1.6.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
=======
	go.mongodb.org/mongo-driver v1.17.6 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/text v0.23.0 // indirect
>>>>>>> origin/main
)
