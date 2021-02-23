# sre-project

An example of a webserver frontend with a Redis backend

Code and configuration files may have verbose comments to explain what happens at each step

This project should work on linux or MacOS (and perhaps Windows if setup properly)

## Prerequisites

- `docker` must be installed
- `docker-compose` must be installed
- `make` must be usable at the command line
- `go` can be instlled but is not required (can run entirely within containers)

## How does it work?

Webserver application publishes on tcp port 9001 and interacts with redis on the backend.

Using docker-compose, this brings a set of docker containers online and the host running them can
open a web browser (or curl) to query http://localhost:9001 and see the Content String set within
the docker-compose.

## Usage

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
