# dkr

## What is dkr?
Dkr is a command line (CLI) tool that is a drop-in replacement for `docker` and `docker-compose` tools.

## What it offers?
Let's say docker CLI interface doesn't provide the best user experience in the world. There are many things to remember, and many traps. Especially for occasional users like me.

This tool adds some handy shortcuts, like for example, instead of writing `docker exec -it CONTAINER_NAME /bin/sh` you can write `dkr sh CONTAINER_NAME`. Quick, right?

## How it works?
By default, it routes all commands directly to docker cli. For example `dkr network ls` would directly call `docker network ls` in your command line.

For convenience it also calls `docker-compose` when explicitly told by calling `dkr c ...`.

If you write this: `dkr i dont know what Im doing` it will evaluate to `docker i dont know what Im doing`.

If `dkr` recognizes known aliases, it would act accordingly as shown in the `dkr sh CONTAINER_NAME` example.

## Aliases
### sh
Enters `CONTAINER_NAME` sh shell.
```shell script
dkr sh CONTAINER_NAME
```
Alias to
```shell script
docker exec -it CONTAINER_NAME /bin/sh
```
### bash
Enters `CONTAINER_NAME` bash shell.
```shell script
dkr bash CONTAINER_NAME
```
Alias to
```shell script
docker exec -it CONTAINER_NAME /bin/bash
```
### killall
Kills all running containers.
```shell script
dkr killall
```
### cleanup
Kills all running containers, removes containers with their images and volumes.
```shell script
dkr cleanup
```

### nuke
Removes everything. Alias to `dkr cleanup` and `docker system prune --volumes -f`.
You may ask here, how come we need to remove running containers, images and volumes first, to properly cleanup the environment with `docker system prune`? 

Well, maintainer of this software doesn't know why, but apparently calling `docker system prune --volumes` when all containers are killed, still leaves some leftovers in the system.

Which brings us to the [primary motto](#how-it-works) of this tool.
```shell script
dkr nuke
```

## Current state
This software is not ready for usage. It's an early stage. Please wait until 1.0.0.

## How to install?
### MacOS via homebrew
First, add ohdkr tap 
```shell script
brew tap ohdkr/homebrew-dkr-osx
```

Then, install `dkr`
```shell script
brew install dkr
```

## How to contribute?
Feel free to post any Issue here on github as bug report, comment or suggestion. Pull requests are always welcome, but it might be wise of you to first create an issue so we can talk details.