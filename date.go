// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"reflect"
	"time"

	"github.com/CloudyKit/jet/v5"
)

// Now wraps time.Now.
func Now(jet.Arguments) reflect.Value {
	return reflect.ValueOf(time.Now())
}

// Date wraps time.Format.
func Date(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("date", 2, 2)
	return reflect.ValueOf(args.Get(0).Interface().(time.Time).Format(args.Get(1).String()))
}
