# sre-project

An example of a webserver frontend with a Redis backend

Code and configuration files may have verbose comments to explain what happens at each step

This project was developed on MacOS and confirmed to work on Ubuntu Linux v20.04. It is possible
it would work on Windows if setup properly.

## Prerequisites

- `docker` must be installed ( https://docs.docker.com/get-docker/ )
- `docker-compose` must be installed ( https://docs.docker.com/compose/install/ )
- `make` must be usable at the command line
- `go` can be installed but is not required (can run entirely within containers)

It is recommended that whatever user runs this be added to the Docker user group permissions
( for example: On Ubuntu linux see `usermod` here https://docs.docker.com/engine/install/ubuntu/ )

## How does it work?

Webserver application publishes on tcp port 9001 and interacts with redis on the backend.

Using docker-compose, this brings a set of docker containers online and the host running them can
open a web browser (or curl) to query http://localhost:9001 and see the Content String set within
the docker-compose.

## Usage

To see what happens with each command, see the `Makefile`

```
make test
```

Uses docker-compose to bring up docker containers (an application and a Redis database) and runs a
simple database ping between the two containers. Useful for testing the stack on any given host.
Caching during the build strep is disabled so that a fresh container is built each time.

```
make run
```

Similar to the test, uses docker-compose to bring up docker containers (an application and a
Redis database). The application reads environment variables provided in the docker-compose.yaml
and sets a Redis key, then returns that key value to a running webserver on TCP port 9001.

Use `Ctrl-C` to break from the running webserver process

```
make all
```

Runs `make test` and then `make run`

```
make clean-dock
```

Cleans out all the container images that were generated each time changes were made.
If containers are in use it may be necessary to add `--force` to the underlying commands used.
