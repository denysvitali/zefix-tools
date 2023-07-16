module github.com/denysvitali/documents-processor

go 1.20

require (
	github.com/alexflint/go-arg v1.4.3
	github.com/denysvitali/documents-indexer v0.0.0
	github.com/denysvitali/go-datesfinder v0.0.1
	github.com/denysvitali/sparql-client v0.0.0
	github.com/opensearch-project/opensearch-go v1.1.0
	github.com/sirupsen/logrus v1.9.3
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.2
)

require (
	github.com/alexflint/go-scalar v1.1.0 // indirect
	github.com/almerlucke/go-iban v0.0.0-20220324081643-09bcab81b879 // indirect
	github.com/denysvitali/go-swiss-qr-bill v0.0.0-20230326211735-9c02af35b762 // indirect
	github.com/goodsign/monday v1.0.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

replace github.com/denysvitali/documents-indexer v0.0.0 => ../documents-indexer

replace github.com/denysvitali/sparql-client v0.0.0 => ../sparql-client
