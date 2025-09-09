### Install goose for db migrations

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Create next directories for supporting db persistence

```bash
mkdir -p $HOME/docker/postgres/shared
mkdir -p $HOME/docker/postgres/17.2/data
```

### Run docker container

```bash
docker-compose -f docker/docker-compose.yml up -d
```

### Run DB migrations:

```bash
goose -dir db/migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
```

## Now you are ready to play with postgres

### Do not forget to stop your docker container after playing

```bash
docker-compose -f docker/docker-compose.yml down
```

