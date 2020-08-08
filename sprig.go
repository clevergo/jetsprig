// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import "github.com/CloudyKit/jet/v5"

// FuncMap is a set of functions.
type FuncMap map[string]jet.Func

// AttachTo attaches function to Set.
func (fm FuncMap) AttachTo(set *jet.Set) {
	for name, fn := range fm {
		set.AddGlobalFunc(name, fn)
	}
}

var genericFuncMap = map[string]jet.Func{
	// string
	"join":       Join,
	"title":      Title,
	"trim":       Trim,
	"trimPrefix": TrimPreffix,
	"trimSuffix": TrimSuffix,

	// date
	"now":  Now,
	"date": Date,
	"ago":  Ago,
}

// GenericFuncMap returns generic functions.
func GenericFuncMap() FuncMap {
	m := make(FuncMap, len(genericFuncMap))
	for name, fn := range genericFuncMap {
		m[name] = fn
	}
	return m
}
