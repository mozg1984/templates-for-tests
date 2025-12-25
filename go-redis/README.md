### Create next directories for supporting storage persistence

```bash
mkdir -p $HOME/docker/redis/8.2-alpine/data
```

### Run docker container

```bash
docker-compose -f docker/docker-compose.yml up -d
```

## Now you are ready to play with redis

### Do not forget to stop your docker container after playing

```bash
docker-compose -f docker/docker-compose.yml down
```