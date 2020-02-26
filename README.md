# Beehive Service
A honeypot is a security system used to detect fake, fraudulent, or otherwise unauthorized access. We use honeypots as both a quality management tool and as a fraud prevention system. The beehive service works by storing tasks that have a known answer. Workers are then assigned to a honeypot question instead of a real task. Since the answer is known, we can tell if the worker was correct or not right away without the use of a verifier. Workers who have consistently failed the honey pot can be prevented from further tasking.

For full [API documentation](https://documenter.getpostman.com/view/7517177/SzKSTKzZ)

## Getting started 

### Prerequisities:

- Install Go (On OS X with Homebrew you can just run `brew install go`.)
- [optional to debug] Postman

### Setup the project

Clone the repository with: 

`go get -u github.com/expandorg/beehive`

OR 

create a directory `$GOPATH/src/github/expandorg` and execute: git clone git@github.com:expandorg/beehive.git 

Run the project dependencies (db, etc.) with `make up`

Run the latest migration with `make migrate-latest`

Run the project with `make run`

### Dependencies

We use `dep` to manage our dependencies.

To add a new vendor, use: 

`deps ensure -add DEPENDENCY`

To update vendors for built project, run:

`make update-deps`

## Database

### Add a new migration

```make add-migration name="migration_name"```

For migration names be descriptive and start with verbs: `create_`, `drop_`, `add_`, etc.

This will look at the latest migrated version (1, 2, 3) and creates 2 files with new version:

`2_migration_name.up.sql` and `2_migration_name.down.sql`

### Migrate

You can migrate to latest:

```make migrate-latest```

OR 

You can migrate up and migrate down a version:

```make run-migrations action="goto" version="1"```

When you migrate up, you can see in the `schema_migrations` the last migrated version. When you migrate down, it updates the the version column in `schema_migrations`.

## Tests
```make run-tests```

### Unit tests
We keep all unit tests close to the code and withing the same package. For example, if you want to test the service package, then you would add the tests in that folder marked `package service`.

### Functional

We keep all functional tests in `tests/` folder. Create a new test file for every function. 

## CI / CD
We use Google Cloud for CI/CD:

*note: please don't modify the following files unless you know what you're doing :)*

**cloudbuild.yaml:** this effectively our CI, it run tests on every PR and will ✓ or x.

**cloudbuild.cd.yaml:** this effectively our CD, it run tests, builds and pushes the image to the container registry and deploys to production on every Master commit, so master has to be always clean. 

**k8s.yaml:** this is the kubernetes setup, including workload and service setup. cloudbuild.cd uses this file to deploy.

## How to Contribute

If you're interested in contributing to the dispute service:

 * Start by reading the [Contributing guide](CONTRIBUTING.md).

## License

Dispute service is licensed under the [MPL-2](license) license
