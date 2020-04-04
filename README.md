# dkr

## What is dkr?
Dkr is a command line (CLI) tool that is a drop-in replacement for `docker` and `docker-compose` tools.

## What it does?
By default, it routes all commands directly to docker cli. For example `dkr network ls` would directly call `docker network ls` in your command line.

For convenience it also calls `docker-compose` when explicitly told by calling `dkr c ...`.  

## So, what it offers?
Let's say docker CLI interface doesn't provide the best user experience in the world. There are many things to remember, and many traps. Especially for occasional users like me.

This tool adds some handy shortcuts, like for example, instead of writing `docker exec -it CONTAINER_NAME /bin/sh` you can write `dkr sh CONTAINER_NAME`. Quick, right?

## How it works?
Everything that is passed after `dkr`, what is not recognised as a superset command, is passed directly to docker(-compose) without any modifications. 
If you write this: `dkr i dont know what Im doing` it will evaluate to `docker i dont know what Im doing`.

If `dkr` recognizes known aliases, it would act accordingly as shown in the `dkr sh CONTAINER_NAME` example.

## Aliases
### sh
```shell script
dkr sh CONTAINER_NAME
```
Alias to
```shell script
docker exec -it CONTAINER_NAME /bin/sh
```
### bash
```shell script
dkr bash CONTAINER_NAME
```
Alias to
```shell script
docker exec -it CONTAINER_NAME /bin/bash
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
Feel free to post any Issue here on github as bug report, comment or suggestion. Pull requests are always welcome.