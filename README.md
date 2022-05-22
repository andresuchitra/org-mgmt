# Organization Management

## Dependencies
- Go version 1.17
- Docker desktop (Mac, Windows)

## How to run
### Local
To run locally, please create `.env` in root project directory with following details.

The DB user is by default `postgres`

```
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=asdf
DB_NAME=mydb
DB_PORT=5432
```

Then run:
```
docker-compose up postgresdb
```

This will run PostgreSQL database at port 5432. Only after successful setup, the Go app can run properly.

### Docker
To run via docker, just execute 
```
make run
```

Please note the docker compose log in the terminal, as the GO app usually runs first before PostgreSQL container successfully and completely setup.

Otherwise, you may run `docker-compose up postgresdb` in the first terminal, and run `docker-compose up api` on second terminal.

## Unit test

To do unit test, please run:
```
make unittest
```



