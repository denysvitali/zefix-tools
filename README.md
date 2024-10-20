# zefix-tools

A collection of tools to interact with the data exported by Zefix, the Swiss Central Business Name Index.

## Get the data

Use [sparql-client](https://github.com/denysvitali/sparql-client) to query the data from 
the [Zefix SPARQL endpoint](https://register.ld.admin.ch/sparql) and generate a JSON file.

```bash
go install github.com/denysvitali/sparql-client/cmd/sparql-client@master
sparql-client --endpoint https://register.ld.admin.ch/query \
    --query query.txt \
    -a -D \
    --limit 500 \
    -o zefix.json
```

## Import the data

### Start PostgreSQL

```bash
docker-compose up -d
```

#### Get the container IP

```bash
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker-compose ps -q | head -n1)
```

Once you have the IP, you can add it to `ZEFIX_DSN` in `.env`

### Import the data to postgres

```bash
source .env
export ZEFIX_DSN
# This will take a while:
go run ./cmd/zefix-import --input ./zefix.json
```

## Query the data

```bash
source .env
export ZEFIX_DSN
go run ./cmd/zefix-find "Rega AG"                                                                                           [23:39:51]
```

will result in:

```plain
Name: Rega AG
URI: https://register.ld.admin.ch/zefix/company/144424
```
