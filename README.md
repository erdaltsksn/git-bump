# Git Bump

[![GoDoc](https://godoc.org/github.com/erdaltsksn/git-bump?status.svg)](https://godoc.org/github.com/erdaltsksn/git-bump)
![Go](https://github.com/erdaltsksn/git-bump/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/git-bump)](https://goreportcard.com/report/github.com/erdaltsksn/git-bump)

Git Bump is a Semantic Version Bumper. It uses the semver and git tags to
define and bump the app version.

## Features

- Cross-Platform
- Integrates into Git

## Requirements

- [Git](https://git-scm.com)

## Getting Started

### 1. Install the application.

```sh
brew install erdaltsksn/tap/git-bump
```

### 2. Bump the version using the interactive cli-ui.

```sh
git bump
```

## Installation

### Using Homebrew

```sh
brew install erdaltsksn/tap/git-bump
```

### Using Go Modules

```sh
go get github.com/erdaltsksn/git-bump
```

## Updating / Upgrading

### Using Homebrew

```sh
brew upgrade erdaltsksn/tap/git-bump
```

### Using Go Modules

```sh
go get -u github.com/erdaltsksn/git-bump
```

## Usage

You may find the documentation for [each command](docs/git-bump.md) inside the
[docs](docs) folder.

### Getting Help

```sh
git bump --help
git bump [command] --help
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](docs/CONTRIBUTING.md) for more information.

## Security Policy

If you discover a security vulnerability within this project, please follow our
[Security Policy Guide](docs/SECURITY.md).

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
