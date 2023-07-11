<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>DDDGo
</h1>
<h3>◦ </h3>
<h3>◦ Developed with the software and tools listed below.</h3>

<p align="center">
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style&logo=Docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/GitHub%20Actions-2088FF.svg?style&logo=GitHub-Actions&logoColor=white" alt="GitHub%20Actions" />

<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/Wire-000000.svg?style&logo=Wire&logoColor=white" alt="Wire" />
<img src="https://img.shields.io/badge/Markdown-000000.svg?style&logo=Markdown&logoColor=white" alt="Markdown" />
</p>
<a href="https://codecov.io/gh/program-world-labs/DDDGo" > 
 <img src="https://codecov.io/gh/program-world-labs/DDDGo/branch/master/graph/badge.svg?token=9V22UZ4HPO"/> 
 </a>
<img src="https://img.shields.io/github/languages/top/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub license" />

</div>

---

## 📒 Table of Contents

- [📒 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [⚙️ Features](#️-features)
- [📂 Project Structure](#-project-structure)
- [🧩 Modules](#-modules)
- [🚀 Getting Started](#-getting-started)
  - [📦 Installation](#-installation)
  - [🧪 Running Tests](#-running-tests)
- [🗺 Roadmap](#-roadmap)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [👏 Acknowledgments](#-acknowledgments)

---

## 📍 Overview

This project is an API service designed and implemented using Domain-Driven Design (DDD) principles. DDD is an approach to software development that greatly emphasizes the importance of understanding and modeling the domain. It aims to ease the creation of complex applications by connecting the related pieces of the software into an evolving model.

The API service is structured around business domains, each represented by a bounded context. The bounded context isolates the domain's models, making them specific to the context and independent from others. This isolation allows the model to evolve independently, reducing the risk of changes in one context affecting others.

The service is built with a layered architecture, typical in DDD:

- **Domain Layer**: This is the core of the software, containing business logic and types, which doesn't depend on other layers.
- **Application Layer**: This layer drives the workflow of the application, directing the domain layer, and is kept thin with logic that does not belong in the domain.
- **Infrastructure Layer**: This layer provides generic technical capabilities that support higher layers (message sending, persistence, drawing UI components, and so on).
- **Adapter Layer (or Presentation Layer)**: This layer is responsible for presenting information to the user and interpreting the user's commands.

This DDD approach ensures that the focus is on the core domain and domain logic, basing complex designs on a model, and initiating a creative collaboration between technical and domain experts to iteratively refine a conceptual model that addresses particular domain problems.

The API service is designed to be scalable, maintainable, and organized around the business domain, providing a strong foundation for further development and adjustments as the understanding of the domain evolves.

---

## ⚙️ Features

- **HTTP Support**: The service can handle HTTP requests and responses.
- **Behavior-Driven Development (BDD)**: The service is developed with a focus on behavior, making it more aligned with business requirements and easier to understand.
- **Event Sourcing**: The state of the business objects is determined by a sequence of events, providing a great audit trail and history that can enable various business insights.
- **SQL Transactions**: The service supports SQL transactions to ensure data consistency and integrity.
- **Anti-Breakdown**
- **Anti-Penetration**
- **Anti-Avalanche**

---

## 📂 Project Structure

```bash
repo
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── app
│       └── main.go
├── config
│   ├── config.go
│   └── dev.yml
├── docker-compose.dev.yml
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── img
│   │   ├── example-http-db.png
│   │   ├── layers-1.png
│   │   ├── layers-2.png
│   │   └── logo.svg
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── integration-test
│   ├── Dockerfile
│   └── integration_test.go
├── internal
│   ├── adapter
│   │   └── http
│   │       ├── error.go
│   │       ├── response.go
│   │       └── v1
│   │           ├── role
│   │           │   ├── error.go
│   │           │   ├── request.go
│   │           │   ├── response.go
│   │           │   └── router.go
│   │           ├── router.go
│   │           └── user
│   │               ├── request.go
│   │               ├── response.go
│   │               └── router.go
│   ├── app
│   │   ├── app.go
│   │   ├── migrate.go
│   │   ├── wire.go
│   │   └── wire_gen.go
│   ├── application
│   │   ├── role
│   │   │   ├── dto.go
│   │   │   ├── error.go
│   │   │   ├── interface.go
│   │   │   └── service.go
│   │   └── user
│   │       ├── dto.go
│   │       ├── interface.go
│   │       └── service.go
│   ├── domain
│   │   ├── aggregate
│   │   │   └── order_aggregate.go
│   │   ├── domainerrors
│   │   │   └── error.go
│   │   ├── entity.go
│   │   ├── repository.go
│   │   ├── search_query.go
│   │   └── user
│   │       ├── entity
│   │       │   ├── amount.go
│   │       │   ├── group.go
│   │       │   ├── role.go
│   │       │   ├── user.go
│   │       │   └── wallet.go
│   │       ├── event
│   │       │   └── user_event.go
│   │       └── repository
│   │           ├── amount_repository.go
│   │           ├── group_repository.go
│   │           ├── role_repository.go
│   │           ├── user_repository.go
│   │           └── wallet_repository.go
│   └── infra
│       ├── amount
│       │   └── repo_impl.go
│       ├── datasource
│       │   ├── cache
│       │   │   ├── cache_local_impl.go
│       │   │   ├── cache_redis_impl.go
│       │   │   └── error.go
│       │   ├── interface.go
│       │   └── sql
│       │       ├── crud_sql_impl.go
│       │       ├── error.go
│       │       ├── transaction_event_impl.go
│       │       └── transaction_run_sql_impl.go
│       ├── dto
│       │   ├── amount.go
│       │   ├── error.go
│       │   ├── group.go
│       │   ├── interface.go
│       │   ├── role.go
│       │   ├── user.go
│       │   ├── utils.go
│       │   └── wallet.go
│       ├── group
│       │   └── repo_impl.go
│       ├── repository
│       │   ├── cache_update_repo_impl.go
│       │   ├── crud_repo_impl.go
│       │   ├── error.go
│       │   └── transaction_repo_impl.go
│       ├── role
│       │   └── repo_impl.go
│       ├── user
│       │   └── repo_impl.go
│       └── wallet
│           └── repo_impl.go
├── migrations
│   ├── 20210221023242_migrate_name.down.sql
│   └── 20210221023242_migrate_name.up.sql
├── pkg
│   ├── cache
│   │   ├── local
│   │   │   └── bigcache_cache.go
│   │   └── redis
│   │       ├── options.go
│   │       └── redis_cache.go
│   ├── httpserver
│   │   ├── options.go
│   │   └── server.go
│   ├── operations
│   │   └── stack_driver.go
│   └── pwsql
│       ├── gorm_mock.go
│       ├── gorm_postgres.go
│       ├── interface.go
│       └── options.go
└── tests
    ├── mocks
    │   ├── DataSource_mock.go
    │   ├── Repository_mock.go
    │   ├── role
    │   │   ├── RoleRepository_mock.go
    │   │   └── RoleService_mock.go
    │   └── user
    │       ├── UserRepository_mock.go
    │       └── UserService_mock.go
    ├── role
    │   ├── create_test.go
    │   ├── features
    │   │   └── usecase
    │   │       ├── role_assigned.feature
    │   │       ├── role_created.feature
    │   │       ├── role_deleted.feature
    │   │       ├── role_list_got.feature
    │   │       └── role_updated.feature
    │   └── sql_test.go
    ├── user
    │   ├── features
    │   │   └── user.feature
    │   └── register_test.go
    └── utils.go

51 directories, 110 files
```

---

## 🧩 Modules

<details closed><summary>Root</summary>

| File                                                                           | Summary                                                                                                                                                                                                                                                                          |
| ------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [go.mod](https://github.com/program-world-labs/DDDGo/blob/main/go.mod)         | The code snippet is a Go module file that lists the required dependencies for a project. It includes various packages and libraries from different sources that are necessary for the project's functionality.                                                                   |
| [Dockerfile](https://github.com/program-world-labs/DDDGo/blob/main/Dockerfile) | The provided code snippet uses Docker to build a Go application in three steps. First, it caches the Go modules. Then, it builds the application. Finally, it creates a minimal Docker image with the necessary files and runs the application.                                  |
| [Makefile](https://github.com/program-world-labs/DDDGo/blob/main/Makefile)     | This code snippet is a Makefile that provides various commands for managing and running a Go application. It includes functionality for running Docker-compose, generating Swagger documentation, running tests, running linters, creating migrations, and mocking dependencies. |

</details>

<details closed><summary>App</summary>

| File                                                                                          | Summary                                                                                                                                                                                                                                                                    |
| --------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [main.go](https://github.com/program-world-labs/DDDGo/blob/main/cmd/app/main.go)              | This code snippet sets up the main application by handling the configuration and running the app with the provided configuration.                                                                                                                                          |
| [wire_gen.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/app/wire_gen.go) | This code snippet sets up dependency injection and wiring for an application. It provides functions to create various services and repositories for users, roles, transactions, and caches. It also initializes an HTTP server with a specified configuration.             |
| [app.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/app/app.go)           | This code snippet initializes a logger and an HTTP server, and then waits for a signal or an error. Upon receiving a signal or error, it shuts down the server.                                                                                                            |
| [wire.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/app/wire.go)         | The provided code snippet is responsible for setting up and providing the necessary dependencies for an HTTP server. It configures database connections, cache sources, repositories, and services. Additionally, it creates an HTTP server with a provided configuration. |
| [migrate.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/app/migrate.go)   | The code snippet initializes a migration process for a PostgreSQL database. It connects to the database, applies pending migrations, and logs the success or failure of the migration. It retries the connection a specified number of times before giving up.             |

</details>

<details closed><summary>Migrations</summary>

| File                                                                                                                                          | Summary                                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [20210221023242_migrate_name.up.sql](https://github.com/program-world-labs/DDDGo/blob/main/migrations/20210221023242_migrate_name.up.sql)     | The code snippet creates a table named "history" with five columns: id, source, destination, original, and translation. The table will be created only if it doesn't already exist. The "id" column is a serial primary key, and the other columns are of type VARCHAR with a maximum length of 255 characters. |
| [20210221023242_migrate_name.down.sql](https://github.com/program-world-labs/DDDGo/blob/main/migrations/20210221023242_migrate_name.down.sql) | The code snippet drops the "history" table if it already exists. This ensures a clean slate for creating or re-creating the table in the future.                                                                                                                                                                |

</details>

<details closed><summary>Config</summary>

| File                                                                                | Summary                                                                                                                                                                                                                                                                                                                   |
| ----------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [config.go](https://github.com/program-world-labs/DDDGo/blob/main/config/config.go) | This code snippet defines a Config struct and a NewConfig function that reads configuration values from environment variables and a YAML file using the viper package. It sets the values of various fields in the Config struct based on the environment and loads the necessary environment variables for GCP services. |

</details>

<details closed><summary>Dto</summary>

| File                                                                                                  | Summary                                                                                                                                                                                                                                                                                                                        |
| ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/error.go)         | This code snippet defines functions for creating error information objects for different DTO transformations, such as user, role, group, wallet, and amount. These functions utilize the domainerrors package to generate unique error codes based on the corresponding DTO and domain errors.                                 |
| [amount.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/amount.go)       | This code snippet defines a data transfer object (DTO) named "Amount" with various methods for transforming, copying, and interacting with the data. It also includes functionality for JSON encoding and decoding, as well as hooks for updating and creating records in a database.                                          |
| [user.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/user.go)           | This code snippet defines a User struct with various properties and methods for data transformation, JSON encoding and decoding, and database operations using GORM. It also includes implementation for interface methods and hooks for updating and creating records.                                                        |
| [role.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/role.go)           | The code snippet defines a Role struct with various fields and methods. It implements the IRepoEntity interface and provides functions for data transformation, conversion to JSON, and decoding JSON. It also includes hooks for updating and creating records in a database, as well as methods for getting and setting IDs. |
| [interface.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/interface.go) | The code snippet defines an interface `IRepoEntity` that extends another interface `IEntity`. It includes methods for retrieving the table name, transforming the entity, converting it to JSON, and decoding JSON back into the entity.                                                                                       |
| [group.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/group.go)         | This code snippet defines a Group struct with various methods for data transformation, database operations, JSON encoding/decoding, and managing timestamps. It also implements interfaces related to repository entities.                                                                                                     |
| [utils.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/utils.go)         | This code snippet defines a function called generateID that generates a random ID of a specified length (10 characters). It utilizes the crypto/rand package to generate random bytes and then converts them to a hexadecimal string representation.                                                                           |
| [wallet.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/dto/wallet.go)       | This code snippet defines a Wallet struct with various attributes and methods. It includes functionality to transform the struct to and from a domain entity, perform database operations, generate UUIDs, serialize to JSON, and decode from JSON.                                                                            |

</details>

<details closed><summary>Datasource</summary>

| File                                                                                                         | Summary                                                                                                                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [interface.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/interface.go) | The provided code snippet defines interfaces for a data source and cache data source, as well as an interface for running transactions. These interfaces define core functionalities such as creating, deleting, updating, and retrieving data, both with and without transactional support. |

</details>

<details closed><summary>Cache</summary>

| File                                                                                                                             | Summary                                                                                                                                                                                                                                                                                   |
| -------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [cache_local_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/cache/cache_local_impl.go) | The provided code snippet is an implementation of a cache data source using the BigCache library. It includes functions for retrieving, storing, and deleting data in the cache using a key generated from the model's table name and ID.                                                 |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/cache/error.go)                       | The code snippet defines error handling functions for cache operations, including setting, deleting, and getting data from the cache. These functions generate error information using an error code and the associated error message.                                                    |
| [cache_redis_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/cache/cache_redis_impl.go) | The provided code snippet is a Go package that implements a Redis cache data source. It includes functions to get data from the cache, set data in the cache, and delete data from the cache. The cache uses a Redis client and interacts with an SQL data source to fetch or store data. |

</details>

<details closed><summary>Sql</summary>

| File                                                                                                                                           | Summary                                                                                                                                                                                                                                                                                                                         |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [transaction_event_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/sql/transaction_event_impl.go)     | The code snippet defines a package for SQL operations. It implements an interface for transaction events using GORM as the ORM library. It provides a method to retrieve the database transaction.                                                                                                                              |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/sql/error.go)                                       | This code snippet defines error handling functions for common SQL operations such as create, delete, update, get, and cast. It uses domain specific error codes to provide more meaningful error messages.                                                                                                                      |
| [crud_sql_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/sql/crud_sql_impl.go)                       | This code snippet is a CRUD (Create, Read, Update, Delete) implementation for interacting with a SQL database using GORM. It provides methods for performing these operations on a given model, both within and outside of a transaction.                                                                                       |
| [transaction_run_sql_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/datasource/sql/transaction_run_sql_impl.go) | The provided code snippet is a part of a SQL package that provides functionality for running transactions using GORM. It includes the implementation of a transaction data source and the ability to run transactions by passing a context and a transaction event function. The code relies on a SQL GORM database connection. |

</details>

<details closed><summary>Repository</summary>

| File                                                                                                                                   | Summary                                                                                                                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/repository/error.go)                                   | This code snippet defines error codes and provides a function to create a new error instance. It handles errors from a datasource and converts them into domain-specific errors with appropriate error codes.                                                                                              |
| [crud_repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/repository/crud_repo_impl.go)                 | The provided code snippet is a CRUD repository implementation that interacts with a database, Redis cache, and local cache. It provides functions for retrieving, creating, updating, and deleting entities. It also includes transactional versions of these functions for use with transactional events. |
| [cache_update_repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/repository/cache_update_repo_impl.go) | This code snippet defines a struct called CacheUpdateImpl that implements methods for saving and deleting data from a remote cache and a local cache. It uses a DTOEntity to transform the data and interacts with the cache through the RemoteCache and Cache data sources.                               |
| [transaction_repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/repository/transaction_repo_impl.go)   | The code snippet provides a TransactionRunRepoImpl implementation that utilizes a TransactionRun datasource to run transactions. It implements the RunTransaction method from the ITransactionRepo interface.                                                                                              |
| [amount_repository.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/repository/amount_repository.go)     | The code snippet defines an interface "AmountRepository" that serves as a repository for handling CRUD operations on the "Amount" domain entity.                                                                                                                                                           |
| [wallet_repository.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/repository/wallet_repository.go)     | The code snippet defines the WalletRepository interface, which inherits from the ICRUDRepository interface. This interface represents a repository responsible for CRUD operations (create, read, update, delete) on wallet domain objects.                                                                |
| [user_repository.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/repository/user_repository.go)         | The code snippet defines a UserRepository interface that extends the ICRUDRepository interface from the domain package. This interface specifies the core functionalities for interacting with user data in a repository.                                                                                  |
| [group_repository.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/repository/group_repository.go)       | The code snippet defines an interface called GroupRepository that inherits methods from another interface ICRUDRepository. This interface is likely to be used for implementing CRUD operations related to group entities in the application's domain.                                                     |
| [role_repository.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/repository/role_repository.go)         | The code snippet defines the RoleRepository interface, which extends the domain.ICRUDRepository interface. This allows it to provide the core functionalities of creating, reading, updating, and deleting roles in the application's domain.                                                              |

</details>

<details closed><summary>Role</summary>

| File                                                                                                           | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/role/repo_impl.go)         | The provided code snippet is a package named "role" that implements the RoleRepository interface. It includes dependencies for user repository, database, cache, and DTO. The RepoImpl struct provides CRUD functionality for roles using these dependencies. The NewRepoImpl function initializes a new RepoImpl instance.                                                                                                                                |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/role/error.go)       | This code snippet includes functions for creating and handling various errors related to the role domain. It also defines a struct for an error event that can be logged using zerolog.                                                                                                                                                                                                                                                                    |
| [response.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/role/response.go) | The code snippet defines a response struct and a function to create a new response object. The response object contains various fields such as ID, name, description, permissions, users, created at, and updated at. The NewResponse function initializes a response object by mapping values from the input model object and creating a list of user responses.                                                                                          |
| [request.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/role/request.go)   | The code defines a struct for a role creation request, including properties for the role's name, description, and permissions. JSON tags and binding requirements are specified to ensure proper data validation and formatting.                                                                                                                                                                                                                           |
| [router.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/role/router.go)     | The provided code snippet is a part of a package for handling role-related functionality in a larger application. It includes a roleRoutes struct and associated methods for creating roles. The create method handles the creation of a role by validating the request parameters, executing the necessary use case, and returning a response. The code also includes imports for various external libraries and packages used in the role functionality. |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/application/role/error.go)           | This code snippet defines error codes and functions for creating specific error types related to domain logic and validation. It allows for consistent error handling and customization within the application.                                                                                                                                                                                                                                            |
| [service.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/application/role/service.go)       | The provided code snippet is a package that defines a service for creating roles. It uses input validation, repository persistence, and logging functionalities.                                                                                                                                                                                                                                                                                           |
| [interface.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/application/role/interface.go)   | The code snippet defines an interface called IService with methods for creating roles. It also includes commented out methods for assigning, updating, and deleting roles, as well as querying for role lists.                                                                                                                                                                                                                                             |
| [dto.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/application/role/dto.go)               | The code snippet provides functionalities for handling roles in an application, including creating and validating role inputs, converting input to entity, handling validation errors, and creating role outputs. It also includes input structs for assigning, updating, and getting lists of roles.                                                                                                                                                      |

</details>

<details closed><summary>Group</summary>

| File                                                                                                    | Summary                                                                                                                                                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/group/repo_impl.go) | The provided code snippet is a Go package that includes a GroupRepository implementation. It uses various dependencies for data storage and caching, and extends a base CRUD implementation. The NewRepoImpl function initializes the repository with the necessary data sources and returns an instance of RepoImpl. |

</details>

<details closed><summary>Amount</summary>

| File                                                                                                     | Summary                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/amount/repo_impl.go) | The code snippet provides a repository implementation for managing amount-related data in the domain. It utilizes CRUD operations and relies on different data sources such as a database and cache to store and retrieve amount data. |

</details>

<details closed><summary>User</summary>

| File                                                                                                           | Summary                                                                                                                                                                                                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/user/repo_impl.go)         | The code snippet defines a user repository implementation that incorporates CRUD operations using a database and cache. It utilizes data transfer objects (DTOs) and is based on an underlying CRUD implementation.                                                                                                                                                            |
| [response.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/user/response.go) | The code snippet defines a Response struct with fields for user information. It provides a NewResponse function that creates a Response object from an application_user.Output model. Only specific fields from the model are included in the Response object.                                                                                                                 |
| [request.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/user/request.go)   | The provided code snippet is a package called "user" that likely contains functionalities related to managing user information or interactions.                                                                                                                                                                                                                                |
| [router.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/user/router.go)     | The code snippet defines a userRoutes struct that handles user-related routes in a web application. It uses the gin framework for routing and the pwlogger package for logging. It includes two routes: getInfo and register, which are not implemented yet. The register function is currently commented out and contains some code related to handling translation requests. |
| [service.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/application/user/service.go)       | The code snippet is a implementation of a user service. It provides functionality to register a user, check if a user already exists, retrieve a user by ID, and handle error cases. It relies on a user repository and a logging interface.                                                                                                                                   |
| [interface.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/application/user/interface.go)   | The code snippet defines an interface called IService that consists of two sets of methods: RegisterUseCase for registering a user and GetByIDUseCase for retrieving a user by ID. These methods take context and user-related information as inputs and return outputs or errors.                                                                                             |
| [dto.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/application/user/dto.go)               | The code snippet defines a struct called Output that represents the output of a user entity. It has properties such as ID, username, email, display name, and avatar. The NewOutput function creates a new Output instance based on a given user entity.                                                                                                                       |

</details>

<details closed><summary>Wallet</summary>

| File                                                                                                     | Summary                                                                                                                                                                                                                           |
| -------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [repo_impl.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/infra/wallet/repo_impl.go) | The code snippet defines a WalletRepository implementation that leverages a CRUDImpl to interact with a database, Redis cache, and another cache. It also provides a constructor function for creating instances of the RepoImpl. |

</details>

<details closed><summary>Http</summary>

| File                                                                                                   | Summary                                                                                                                                                                                                                                                                                                                     |
| ------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/error.go)       | The code snippet defines an AdapterError struct that wraps around a domain error. It also provides a function to create a new AdapterError instance by converting an error into a domain error.                                                                                                                             |
| [response.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/response.go) | The code snippet provides functionalities for handling error responses and generating success responses in an HTTP server using the Gin framework. It includes a struct for the response format, functions for handling error responses and success responses, and an adapter for converting errors to the response format. |

</details>

<details closed><summary>V1</summary>

| File                                                                                                  | Summary                                                                                                                                                                                                                                                                                |
| ----------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [router.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/adapter/http/v1/router.go) | The code snippet implements routing paths for a Go API. It includes functionality for handling Swagger documentation, Kubernetes probes, Prometheus metrics, and routes for user and role operations. The code uses the Gin framework and follows a clean architecture design pattern. |

</details>

<details closed><summary>Domain</summary>

| File                                                                                                     | Summary                                                                                                                                                                                                                                                                                                                                                                                                                     |
| -------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [search_query.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/search_query.go) | The code snippet defines a package "domain" containing struct types and methods for a search query feature. The "Page" struct represents pagination information, the "Filter" struct represents query filters, and the "Sort" struct represents sorting criteria. The "SearchQuery" struct combines these elements. The code also includes methods to retrieve the WHERE clause, arguments, and order for the search query. |
| [entity.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/entity.go)             | The code snippet defines an interface called IEntity that includes methods for getting and setting the ID of an entity.                                                                                                                                                                                                                                                                                                     |
| [repository.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/repository.go)     | The code snippet defines interfaces for CRUD operations, transaction handling, and cache updates in a domain package. It provides methods for retrieving, creating, updating, and deleting entities, as well as transactional versions of these operations. Additionally, there is an interface for cache-related operations.                                                                                               |

</details>

<details closed><summary>Domainerrors</summary>

| File                                                                                                    | Summary                                                                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [error.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/domainerrors/error.go) | The provided code snippet defines a package for handling domain-specific errors. It includes functions for creating new errors, checking error codes, and retrieving the underlying cause of an error. The code also includes a struct for storing error information and implements the error interface. |

</details>

<details closed><summary>Aggregate</summary>

| File                                                                                                                     | Summary                                                                                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [order_aggregate.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/aggregate/order_aggregate.go) | The code snippet is part of a package called "aggregate". It likely contains functions or classes that perform aggregation operations on data. Further details or specific functionality cannot be determined without examining the rest of the code. |

</details>

<details closed><summary>Entity</summary>

| File                                                                                                     | Summary                                                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [amount.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/entity/amount.go) | The code snippet defines an Amount struct that represents a monetary amount. It implements the IEntity interface with GetID and SetID methods. It includes fields for ID, currency, icon, balance, decimal, wallet ID, creation/update/delete timestamps.                                                        |
| [user.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/entity/user.go)     | The provided code snippet defines a User entity structure with various fields such as ID, Username, Password, and more. It implements the IEntity interface and provides methods to get and set the ID. There is also a NewUser function to create a new User instance.                                          |
| [role.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/entity/role.go)     | This code snippet defines a Role struct with various attributes such as ID, name, description, permissions, users, and timestamps. It also includes methods to get and set the ID. The NewRole function creates a new Role instance with specified attributes. The code implements the domain.IEntity interface. |
| [group.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/entity/group.go)   | This code snippet defines a Group entity with properties like ID, Name, Description, Users, Owner, Metadata, CreatedAt, UpdatedAt, and DeletedAt. It also implements the domain.IEntity interface methods for getting and setting the ID.                                                                        |
| [wallet.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/entity/wallet.go) | The provided code snippet defines a Wallet struct with various properties such as ID, Name, Description, Chain, Address, UserID, Amounts, CreatedAt, UpdatedAt, and DeletedAt. It also implements the IEntity interface and includes methods for getting and setting the ID of the wallet.                       |

</details>

<details closed><summary>Event</summary>

| File                                                                                                            | Summary                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [user_event.go](https://github.com/program-world-labs/DDDGo/blob/main/internal/domain/user/event/user_event.go) | The code snippet defines a package called "event." It is likely to contain functionality related to events, such as event creation, handling, and management. |

</details>

<details closed><summary>Integration-test</summary>

| File                                                                                                              | Summary                                                                                                                                                                                                                                                                                                                                                                                          |
| ----------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [Dockerfile](https://github.com/program-world-labs/DDDGo/blob/main/integration-test/Dockerfile)                   | This code snippet sets up a multi-stage build process. In the first stage, it caches the Go modules by copying the go.mod and go.sum files, downloading the modules using'go mod download'. In the second stage, it copies the downloaded modules, source code, and sets the necessary environment variables for building a Linux binary. Finally, it runs the integration tests using'go test'. |
| [integration_test.go](https://github.com/program-world-labs/DDDGo/blob/main/integration-test/integration_test.go) | The provided code snippet is an integration test that checks the availability of a host by performing a health check HTTP request. It retries the request a specified number of times until a successful response is received or the attempts are exhausted. If the host is not available, it logs an error and exits the test.                                                                  |

</details>

<details closed><summary>Redis</summary>

| File                                                                                                   | Summary                                                                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [options.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/cache/redis/options.go)         | The code snippet provides a package for interacting with Redis. It includes two options: MaxRetries sets the maximum number of retries for failed operations, and RetryDelay sets the delay between retries.                                              |
| [redis_cache.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/cache/redis/redis_cache.go) | This code snippet provides a Redis struct with a New function for creating a Redis client using the provided Redis URL. It supports custom options for max retries and retry delay. The Close function is available to gracefully close the Redis client. |

</details>

<details closed><summary>Local</summary>

| File                                                                                                         | Summary                                                                                                                                                                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [bigcache_cache.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/cache/local/bigcache_cache.go) | The provided code snippet is a wrapper around the BigCache library, which helps in creating and managing a high-performance, in-memory cache. It sets up the cache with various configurations, such as the number of shards, expiration time, memory allocation, size limit, and callbacks. It also provides functions to create and close the cache instance. |

</details>

<details closed><summary>Pwsql</summary>

| File                                                                                                 | Summary                                                                                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [options.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/pwsql/options.go)             | The code snippet defines three options (MaxPoolSize, ConnAttempts, and ConnTimeout) for configuring a Postgres database connection. These options modify the Postgres struct to set the maximum pool size, connection attempts, and connection timeout respectively.   |
| [gorm_postgres.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/pwsql/gorm_postgres.go) | The provided code snippet is a package for interacting with a PostgreSQL database using the Gorm ORM. It establishes a connection to the database, sets pool size and connection attempts, and provides methods for accessing the database and closing the connection. |
| [interface.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/pwsql/interface.go)         | The code snippet provides an interface named ISQLGorm with two core functionalities-GetDB() returns an instance of gorm.DB, and Close() closes the connection to the database.                                                                                         |
| [gorm_mock.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/pwsql/gorm_mock.go)         | The code snippet provides a package "pwsql" for working with a SQL database using GORM. It includes functions for initializing, getting the database connection, and closing the connection.                                                                           |

</details>

<details closed><summary>Operations</summary>

| File                                                                                                    | Summary                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [stack_driver.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/operations/stack_driver.go) | The code snippet initializes and configures OpenTelemetry for Google Cloud Platform (GCP) operations, including trace exporting, resource detection, and setting the tracer provider. |

</details>

<details closed><summary>Httpserver</summary>

| File                                                                                          | Summary                                                                                                                                                                                                                                                                                                                  |
| --------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [server.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/httpserver/server.go)   | The code snippet defines a package that implements an HTTP server. It allows creating and starting an HTTP server with custom options. It also provides methods for receiving notifications and gracefully shutting down the server. The default settings for timeouts, address, and shutdown duration are also defined. |
| [options.go](https://github.com/program-world-labs/DDDGo/blob/main/pkg/httpserver/options.go) | This code snippet provides options to configure a HTTP server. It allows setting the server's port, read and write timeouts, and shutdown timeout. These options can be customized using the provided functions.                                                                                                         |

</details>

---

## 🚀 Getting Started

### 📦 Installation

1. Clone the DDDGo repository:

```sh
git clone https://github.com/program-world-labs/DDDGo
```

2. Change to the project directory:

```sh
cd DDDGo
```

3. Install the dependencies:

```sh
go mod tidy
```

### 🧪 Running Tests

```sh
make test
```

---

## 🗺 Roadmap

> - [x] `ℹ️ Role Implementation`
> - [ ] `ℹ️ Event Sourcing Implementation`
> - [ ] `ℹ️ User Implementation`
> - [ ] `ℹ️ Group Implementation`
> - [ ] `ℹ️ Wallet Implementation`

---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:

1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).

```sh
git checkout -b new-feature-branch
```

4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.

```sh
git commit -m 'Implemented new feature.'
```

6. Push your changes to your forked repository on GitHub using the following command

```sh
git push origin new-feature-branch
```

7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
   The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `ℹ️ INSERT-LICENSE-TYPE` License. See the [LICENSE](https://docs.github.com/en/communities/setting-up-your-project-for-healthy-contributions/adding-a-license-to-a-repository) file for additional info.

---

## 👏 Acknowledgments

---
