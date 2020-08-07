// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"reflect"
	"strings"

	"github.com/CloudyKit/jet/v5"
)

// Title wraps strings.Title.
func Title(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("title", 1, 1)
	return reflect.ValueOf(strings.Title(args.Get(0).String()))
}

// Trim wraps strings.TrimSpace.
func Trim(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("trim", 1, 1)
	return reflect.ValueOf(strings.TrimSpace(args.Get(0).String()))
}

// TrimPreffix wraps strings.TrimPreffix.
func TrimPreffix(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("trimPrefix", 2, 2)
	return reflect.ValueOf(strings.TrimPrefix(args.Get(0).String(), args.Get(1).String()))
}

// TrimSuffix wraps strings.TrimSuffix.
func TrimSuffix(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("trimPSuffix", 2, 2)
	return reflect.ValueOf(strings.TrimSuffix(args.Get(0).String(), args.Get(1).String()))
}

// Join wraps strings.Join.
func Join(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("join", 2, 2)
	return reflect.ValueOf(strings.Join(args.Get(0).Interface().([]string), args.Get(1).String()))
}
