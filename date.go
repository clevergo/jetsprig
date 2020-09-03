// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"reflect"
	"time"

	"github.com/CloudyKit/jet/v5"
)

// Now returns the current local time.
func Now(jet.Arguments) reflect.Value {
	return reflect.ValueOf(time.Now())
}

// Date formats a date with the given layout.
func Date(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("date", 1, 2)
	format := "2006-01-02"
	if args.NumOfArguments() > 1 {
		format = args.Get(1).String()
	}
	return reflect.ValueOf(args.Get(0).Interface().(time.Time).Format(format))
}

// DateInZone returns the copy of same time instant with the given
// time zone.
func DateInZone(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("dateInZone", 2, 2)
	loc, err := time.LoadLocation(args.Get(1).String())
	if err != nil {
		panic(err)
	}
	return reflect.ValueOf(args.Get(0).Interface().(time.Time).In(loc))
}

// Ago returns duration from time.Now in seconds resolution.
func Ago(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("ago", 1, 1)
	u := args.Get(0).Interface().(time.Time)
	return reflect.ValueOf(time.Now().Sub(u).String())
}
