// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import "github.com/CloudyKit/jet/v5"

var funcMap = map[string]jet.Func{
	// string
	"join":       Join,
	"title":      Title,
	"trim":       Trim,
	"trimPrefix": TrimPreffix,
	"trimSuffix": TrimSuffix,

	// date
	"now": Now,
}

// FuncMap returns all functions.
func FuncMap() map[string]jet.Func {
	m := make(map[string]jet.Func, len(funcMap))
	for name, fn := range funcMap {
		m[name] = fn
	}
	return m
}

// AttachTo attaches functions to Set.
func AttachTo(funcMap map[string]jet.Func, set *jet.Set) {
	for name, fn := range funcMap {
		set.AddGlobalFunc(name, fn)
	}
}
