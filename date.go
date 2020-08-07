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
