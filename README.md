# Jet template engine functions
[![Build Status](https://img.shields.io/travis/clevergo/jetsprig?style=for-the-badge)](https://travis-ci.org/clevergo/jetsprig)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/jetsprig?style=for-the-badge)](https://coveralls.io/github/clevergo/jetsprig)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/jetsprig?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jetsprig?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/jetsprig)
[![Release](https://img.shields.io/github/release/clevergo/jetsprig.svg?style=for-the-badge)](https://github.com/clevergo/jetsprig/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/month/clevergo.tech/jetsprig&style=for-the-badge)](https://pkg.clevergo.tech/)

```shell
$ go get -u clevergo.tech/jetsprig
```

## Usage

```go
import (
    "clevergo.tech/jetsprig"
    "github.com/CloudyKit/jet/v5"
)

set := jet.NewHTMLSet("")
// attaches functions to Set.
jetsprig.AttachTo(jetsprig.FuncMap(), testSet)
```

## Functions

| **String** | |
|:---|:---
| `trim` | `strings.TrimSpace` 
| `trimPrefix` | `strings.TrimPrefix` 
| `trimSuffix` | `strings.TrimSuffix` 
| **Date** 
| `now` | `time.Now` 