### Install goose for db migrations

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Create next directories for supporting db persistence

```bash
mkdir -p $HOME/docker/clickhouse/25.4/data
mkdir -p $HOME/docker/clickhouse/25.4/logs
```

### Run docker container

```bash
docker-compose -f docker/docker-compose.yml up -d
```

### Run DB migrations:

```bash
goose -dir db/migrations clickhouse "tcp://$CLICKHOUSE_USERNAME:$CLICKHOUSE_PASSWORD@$CLICKHOUSE_HOST:$CLICKHOUSE_PORT/$CLICKHOUSE_DATABASE?secure=$CLICKHOUSE_SSL" up
```

## Now you are ready to play with postgres

### Do not forget to stop your docker container after playing

```bash
docker-compose -f docker/docker-compose.yml down
```