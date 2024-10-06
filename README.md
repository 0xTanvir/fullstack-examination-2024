# Todo APP - Fullstack Technical Challenge

## Refactoring
I refactored the project into distinct layers (`api`, `domain`, `infrastructure`) to improve code separation, maintainability, and testability:
- api: Contains the input/output logic of the application.
- domain: Holds the business logic and entities of the application.
- infrastructure: Contains the implementation details, such as data storage and configuration.

I also removed `common` and `model` packages, which I found too vague and potentially problematic for long-term organization as the project grows.
- `common` was removed, and its contents, like `version` and `config`, were moved to `infrastructure`. They could be relocated to `project-root/pkg/` if they are shared across packages.
- `model` was refactored into `dto` (Data Transfer Objects) and `bo` (Business Objects) for better organization by functionality.
- All interface definitions were consolidated into a single location (`definition`).
- Duplicated structures like `DeleteRequest` and `FindRequest` were combined into one (`IDWrapper`) in `api/dto/requests.go`.

Project structure after refactoring:

```
internal
├── api
│   ├── dto
│   │   ├── ...
│   │   └── todo.go
│   ├── handler
│   │   ├── ...
│   │   └── validator.go
│   └── server
│       ├── ...
│       └── server.go
├── domain
│   ├── bo
│   │   ├── ...
│   │   └── todo.go
│   ├── definition
│   │   ├── ...
│   │   └── datastore.go
│   └── service
│       └── todo.go
└── infrastructure
    ├── config
    │   └── server.go
    ├── datastores
    │   └── sqlite
    │       ├── ...
    │       └── todo.go
    └── version
        ├── ...
        └── version.go
```

## Code Formatting
- Followed Go best practices with consistent indentation, line length, spacing, and naming conventions (camelCase).
- Grouped related code together and ensured proper commenting, adhering to golangci-lint standards.

## Documentation
- Fixed a Swagger documentation issue where `/:id` was changed to `{id}`.
- Updated Swagger documentation to reflect all the API changes. The documentation is now fully functional and can be used to test the API.
- Used the Conventional Commits standard for Git commit messages.

## Question 1
- The `/api/v1/todos` endpoint now accepts filters like `task` and `status`, as well as sorting and ordering by `priority`.
- A new search field and sort button have been added to the UI.
- Swagger API documentation has been updated for these changes.
## Question 2
- Added a `priority` field to the Todo model.
- A "Sort by Priority" button was added to the UI next to the search field.
- The backend API is configured to return higher-priority tasks at the top.
## Question 3
- The UI now has two lists: `Active Tasks` and `Completed Tasks`.
- Tasks marked as `done` move to the `Completed Tasks` list at the bottom.
- Tasks changing from `done` to another status return to the `Active Tasks` list at the top.


# Fullstack examination 2024

## What is Fullstack examination 2024?

This is a repository for exams used to assess technical skills for hiring full-stack engineers (Golang).

## Dev

You will set up the development environment using asdf, a tool for managing multiple language versions.
If you are in an environment where asdf cannot be used, please install the necessary tools according to your environment.

### Install

#### Install asdf

[Getting Started | asdf](https://asdf-vm.com/guide/getting-started.html)

#### Install asdf plugins

```bash
asdf plugin add air
asdf plugin add golang
asdf plugin add golangci-lint
asdf plugin add gotestsum
asdf plugin add nodejs
asdf plugin add swag
```

#### Install asdf versions

Please install the versions listed in the `.tool-versions` file.

```bash
asdf install
```

### Start Development Environment

#### backend

```bash
make dep-backend-local
```

```bash
make migrate
```

```bash
make serve-backend
```

#### ui

```bash
make dep-ui-local
```

```bash
make serve-ui
```

When you access [http://localhost:3000/](http://localhost:3000/), the UI screen will be displayed.

### Migration

Please run the migration process.

```bash
make migrate
```

If the migration fails due to the current state of the schema, please delete the database and run the migration again.

```bash
make reset-local-db migrate
```

### Format

To maintain consistency in the code, formatting should be applied. Be sure to run it once development is complete.

```bash
make fmt
```

### Swagger

Please create API documentation using Swagger.

```bash
make swagger
```

To generate documentation, you need to add comments to your Go files. Please refer to this documentation for the correct syntax and guidelines.

[swaggo/swag: Automatically generate RESTful API documentation with Swagger 2.0 for Go.](https://github.com/swaggo/swag#declarative-comments-format)

While running `make serve-backend`, you can access [http://localhost:1314/](http://localhost:1314/) to view the Swagger UI. Or alternatively, you can open `docs/swagger.yaml` to view the API documentation.

### Test

Please run the tests.

```bash
make test-backend
```

## examination

Please solve the tasks outlined in [examination.md](./examination.md).
