# challenge
This is an authorization service for a coding challenge.

## Setup and Usage
### Setup
1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop).
2. Install [sql-migrate](https://github.com/rubenv/sql-migrate)
3. Make sure you're running Go 1.11.x

#### Windows
1. Open a Powershell terminal and navigate to the project root.
2. Execute `.\start.bat`
3. After that completes, execute `go run main.go`

#### Mac or Linux
I don't have an environment in which to reliably test a shell script. Sorry.
1. Open a terminal
2. From the project root, execute `docker-compose up cockroach`
3. From the project root, execute `cd ./scripts/sql`
3. Execute `sql-migrate up`
4. Execute `cd ../..`
5. Execute `go run main.go`

### Usage
1. Grab the Postman collection under `tools/postman`
2. Log in with the default admin (credentials below)
3. Grab the JWT returned in the response body
4. Make more requests. Subsequent JWTs will be in the response header.

default admin:
- username: `bpechiney`
- password: `letmein`

default admin:
- username: `bpechiney`
- password: `letmein`

## Documentation
Contents:
- API:
- Design
- Limitations
- Packages Used

### API
Request bodies are provided in the Postman collection.
- `POST /api/v1/signup`
- `POST /api/v1/login`
- `GET /api/v1/attestations`
- `POST /api/v1/attestations`
- `POST /api/v1/register-priv`
- `GET /api/v1/users`

### Design
Contents:
- Architecture
- Organization

#### Architecture
The database is a very simple one implemented with [CockroachDB](https://www.cockroachlabs.com/). I chose
CockroachDB because it is based on PostgreSQL - which I am very familiar with - and it was designed and
written with the cloud and redundancy in mind. It also offers clever features like
[interleaved tables](https://www.cockroachlabs.com/docs/v2.1/interleave-in-parent.html#main-content).

The schema consists of two tables: `challenge_user` and `attestation`. `attestation` has two foreign key
references to `challenge_user`: `claimant_id` and `attestor_id`. A common table expression is used to
retrieve attestations.

Routes are protected by a middleware that attempts to inspect the Authorization header of the incoming
request. It responds with an authorization error if that inspection fails.

#### Organization
The project is organized into five packages:
- challenge
- cockroach
- http
- mock

##### challenge
The challenge package contains the application's domain-specific information. With the exception of the
errors package, it references only built-in packages and the packages defined in the application. This provides
a very modular, decoupled implementation that allows for significant refactors with little pain.

##### cockroach
The cockroach package provides a data access object that abstracts the underlying [CockroachDB](https://www.cockroachlabs.com/)
database by wrapping sql.DB and providing a DAO constructor that connects to the database.

##### http
The http package contains the request handlers, routes, and middleware of the application. It handles all
network communication as well as authentication and authorization. It contains a `server` struct which
provides access to the repositories and configuration information the request handlers need.

##### mock
The mock package contains functions that provide mocks for unit tests. It is not used because I ran out of
time and had to skimp on unit tests.

### Limitations
The following are open issues due to time constraint.
#### Security
- communication with the database should be TLS encrypted
- communication with the API should be TLS encrypted
- JWTs should be signed with a robust key that is kept outside of source control (maybe [Vault](https://www.vaultproject.io/))
- database connection information should be injected via a CI/CD pipeline

#### Features
- error responses should be more specific
- database errors should be handled more elegantly and be more specific
- a "forgot password" feature should be added
- request (correlation) IDs should be injected into every request
- timeouts should be implemented via the [context package](https://golang.org/pkg/context/)

#### Code Quality
- unit test coverage should exceed 90%
- an error code system should be implemented
- there should be "generic" request and response structs with common attributes
- there should be structured logging (maybe [logrus](https://github.com/sirupsen/logrus))
- logs should be written to a temporal database
- there should be audit logging
- Swagger docs

### Packages Used
- [viper](https://github.com/spf13/viper) for configuration
- [jwt-go](https://github.com/dgrijalva/jwt-go) for JWT authentication
- [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) for securely handling passwords
- [mux](https://github.com/gorilla/mux) for request routing
- [errors](https://github.com/pkg/errors) for enhanced errors
