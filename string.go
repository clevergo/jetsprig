// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"reflect"
	"strings"

	"github.com/CloudyKit/jet/v5"
)

// Title converts a string to title case.
func Title(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("title", 1, 1)
	return reflect.ValueOf(strings.Title(args.Get(0).String()))
}

// Trim removes space from either side of a string.
func Trim(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("trim", 1, 1)
	return reflect.ValueOf(strings.TrimSpace(args.Get(0).String()))
}

// TrimPrefix removes the prefix of a string.
func TrimPrefix(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("trimPrefix", 2, 2)
	return reflect.ValueOf(strings.TrimPrefix(args.Get(0).String(), args.Get(1).String()))
}

// TrimSuffix removes the suffix of a string.
func TrimSuffix(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("trimPSuffix", 2, 2)
	return reflect.ValueOf(strings.TrimSuffix(args.Get(0).String(), args.Get(1).String()))
}

// Join concatenates the elements of its first argument to create a single string.
func Join(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("join", 2, 2)
	return reflect.ValueOf(strings.Join(args.Get(0).Interface().([]string), args.Get(1).String()))
}

// Abbrev
func Abbrev(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("join", 2, 2)
	s := args.Get(0).String()
	width := args.Get(1).Interface().(int)
	if len(s) < 4 || width < 4 || len(s) < width-3 {
		return reflect.ValueOf(s)
	}
	return reflect.ValueOf(s[:width-3] + "...")
}
