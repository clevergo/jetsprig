# Jet template engine functions
[![Build Status](https://img.shields.io/travis/clevergo/jetsprig?style=flat-square)](https://travis-ci.org/clevergo/jetsprig)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/jetsprig?style=flat-square)](https://coveralls.io/github/clevergo/jetsprig)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/clevergo.tech/jetsprig?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jetsprig?style=flat-square)](https://goreportcard.com/report/github.com/clevergo/jetsprig)
[![Release](https://img.shields.io/github/release/clevergo/jetsprig.svg?style=flat-square)](https://github.com/clevergo/jetsprig/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/total/clevergo.tech/jetsprig&style=flat-square)](https://pkg.clevergo.tech/clevergo.tech/jetsprig)
[![Chat](https://img.shields.io/badge/chat-telegram-blue?style=flat-square)](https://t.me/clevergotech)
[![Community](https://img.shields.io/badge/community-forum-blue?style=flat-square&color=orange)](https://forum.clevergo.tech)

This library inspired by [Masterminds/sprig)](https://github.com/Masterminds/sprig).

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
jetsprig.GenericFuncMap().AttachTo(set)
```

## Functions

| **String**      |                     |
|:----------------|:---
| `trim`          | `strings.TrimSpace` 
| `trimPrefix`    | `strings.TrimPrefix` 
| `trimSuffix`    | `strings.TrimSuffix` 
| `join`          | `strings.Join`
| **Date** 
| `now`           | `time.Now` 
| `date`          | `time.Time.Format`
