<img src="https://leaders2021.innoagency.ru/static/img/general/logo.svg"
  style="height: 80px;">

# Sport Object Analysis

## Depolyment
1. Clone repository and "cd" in it
```bash
$ git clone https://github.com/techpotion/leaders2021-backend.git
$ cd leaders2021-backend
```

2. Fill the config.yaml, according to config.example.yaml example
```yaml
server:
  grpc:
    interface: 0.0.0.0
    port: 3200
  http:
    interface: 0.0.0.0
    port: 3201

microservices:
  marks:
    host: 192.168.1.58
    port: 3300

database:
  host: 192.168.1.58
  port: 5432
  username: postgres
  password: postgres
  db: postgres
```

Alternatively, you may set database connection data using environment variables
```bash
$ export DB_DB=postgis
$ export DB_HOST=localhost
$ export DB_PORT=5432
$ export DB_USERNAME=postgis
$ export DB_PASSWORD=postgis
$ export MARKS_HOST=localhost
$ export MARKS_PORT=3300
$ export EXPORTS_HOST=localhost
$ export EXPORTS_PORT=3400
```
\* Environment variables take precedence over data in config.yaml file

3. Install Golang if you need to - https://golang.org/doc/install

4. Install project's dependencies
```bash
$ go mod download
```

5. Run the project
```bash
$ go run main.go
```

## Depolyment using Docker
1.  Clone repository and "cd" in it
```bash
$ git clone https://github.com/techpotion/leaders2021-backend.git
$ cd leaders2021-backend
```

2. Build the docker image
```bash
$ docker build --tag backend .
```

3. Run the project passing database data as environment variables via `-e` argument
```bash
$ docker run -d -p 3200:3200 -p 3201:3201 --name techpotion-leaders2021-backend -e DB_DB=postgis -e DB_HOST=192.168.1.58 -e DB_PORT=5542 -e DB_USERNAME=postgis -e DB_PASSWORD=postgis backend
```