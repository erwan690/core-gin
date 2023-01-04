# Core Gin

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-80%25-brightgreen.svg?longCache=true&style=flat)</a>

Core Gin with [Gin Web Framework](https://github.com/gin-gonic/gin)

## Features

-   Application backbone with [Gin Web Framework](https://github.com/gin-gonic/gin)
-   Dependency injection using [uber-go/fx](https://pkg.go.dev/go.uber.org/fx)
-   API endpoints documentation with [swag](https://github.com/swaggo/swag)
-   Uses fully featured [GORM](https://gorm.io/index.html)
-   Agent [OpenTelemetry-Go](https://pkg.go.dev/go.opentelemetry.io/otel)
-   Metric [go-http-metrics](https://github.com/slok/go-http-metrics)


## Run application

-   Setup environment variables

```zsh
cp .env.example .env
```

-   Update your database credentials environment variables in `.env` file

### Locally

-   Run `go run main.go app:serve` to start the server.
-   There are other commands available as well. You can run `go run main.go -help` to know about other commands available.


## Folder Structure :file_folder:

| Folder Path                      | Description                                                                                         |
| -------------------------------- | --------------------------------------------------------------------------------------------------- |
| `/api`                           | contains all the `middlwares`, and `routes` of the server in their respective folders               |
| `/bootstrap`                     | contains modules required to start the application                                                  |
| `/commands`                      | server commands                                                                                     |
| `/constants`                     | global application constants                                                                        |
| `/docs`                          | API endpoints documentation using `swagger`                                                         |
| `/infrastructure`                | third-party services connections like `postgress`, `otel`, `slack`, ...                             |
| `/lib`                           | contains library code                                                                               |
| `/migration`                     | database migration files                                                                            |
| `/internal/dto`                  | Data Transfer Object. such as Request and response Model                                            |
| `/internal/handler`              | handler layer it's such as controller to handle routes endpoint                                     |
| `/internal/models`               | ORM models                                                                                          |
| `/internal/repositories`         | contains repository. Mainly database queries are added here.                                        |
| `/internal/services`             | service layers, contains the functionality that compounds the core of the application               |
| `/utils`                         | global utility/helper functions                                                                     |
| `.env.example`                   | sample environment variables                                                                        |
| `dbconfig.yml`                   | database configuration file for `sql-migrate` command                                               |
| `main.go`                        | entry-point of the server                                                                           |
| `Makefile`                       | stores frequently used commands; can be invoked using `make` command                                |

---

## Migration Commands

⚓️ &nbsp; If you want to run the migration runner from the host environment instead of the docker environment; ensure that `sql-migrate` is installed on your local machine.

### Install `sql-migrate`

> You can skip this step if `sql-migrate` has already been installed on your local machine.

**Note:** Starting in Go 1.17, installing executables with `go get` is deprecated. `go install` may be used instead. [Read more](https://go.dev/doc/go-get-install-deprecation)

```zsh
go install github.com/rubenv/sql-migrate/...@latest
```

If you're using Go version below `1.18`

```zsh
go get -v github.com/rubenv/sql-migrate/...
```

### Running migration

<details>
    <summary>Available migration commands</summary>

| Command               | Desc                                                       |
| --------------------- | ---------------------------------------------------------- |
| `make migrate-status` | Show migration status                                      |
| `make migrate-up`     | Migrates the database to the most recent version available |
| `make migrate-down`   | Undo a database migration                                  |
| `make migrate-redo`   | Reapply the last migration                                 |
| `make migrate-create` | Create new migration file                                  |

</details>

---

## Testing

The framework comes with unit and integration testing support out of the box. You can check examples written in tests directory.

To run the test just run:

```zsh
go test ./... -v
```

### For test coverage

```zsh
go test ./... -v -coverprofile cover.txt -coverpkg=./...
go tool cover -html=cover.txt -o index.html
```

### For coverage Badge

> You can skip this step if `gopherbadger` has already been installed on your local machine.

install the executeable (ensure your $PATH contains $GOPATH/bin):

```
go install github.com/jpoles1/gopherbadger@latest
```
### Running generation Bade

<details>
    <summary>Available generate badge commands</summary>

| Command               | Desc                                                       |
| --------------------- | ---------------------------------------------------------- |
| `make generate-cover` | Generate Coverage Badge on README                          |

</details>

---

## API documents with SWAG


> You can skip this step if `swag` has already been installed on your local machine.

- Download swag by using:
```zsh
go install github.com/swaggo/swag/cmd/swag@latest
```

- Add comments to your API source code, See [Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).

### Running generation

<details>
    <summary>Available generate commands</summary>

| Command               | Desc                                                       |
| --------------------- | ---------------------------------------------------------- |
| `make generate-docs`  | Generate Documentation On Declarative Comments Format      |

</details>

To See Doc run the app and Browse to `http://localhost/swagger/index.html`

-   You can see all the documented endpoints in Swagger-UI from the API specification

---