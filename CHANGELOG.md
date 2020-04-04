# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] (0.1.0)
## Added
- [dkr sh CONTAINER_NAME](./README.md#sh) command.
- [dkr bash CONTAINER_NAME](./README.md#bash) command.

## Changed
- A `--version` flag now prints version of `dkr`, `docker` and `docker-compose`.

## Fixed
- Index out of range error when calling `dkr` or `dkr c` without any more arguments.

## [v0.0.1](https://github.com/ohdkr/dkr/releases/tag/v0.0.1)
### Added
- Basic proxy to docker (`dkr [args]`) and docker-compose (`dkr c [args]`).
- A `--version` flag which makes `dkr` output current version.