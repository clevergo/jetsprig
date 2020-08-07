// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"strings"
	"testing"
)

func TestTitle(t *testing.T) {
	for _, s := range []string{"foo", "bar", "foo bar"} {
		executeTestTemplate(t, "title", `{{ .s | title }}`, nil, map[string]string{"s": s}, strings.Title(s))
	}
}

func TestTrim(t *testing.T) {
	for _, s := range []string{"foo", " foo", "foo "} {
		executeTestTemplate(t, "trim", `{{ .s | trim }}`, nil, map[string]string{"s": s}, strings.TrimSpace(s))
	}
}

func TestTrimPrefix(t *testing.T) {
	prefix := "foo"
	for _, s := range []string{"bar", " foobar", "foo bar"} {
		executeTestTemplate(t, "title", `{{ trimPrefix(.s, .prefix) }}`, nil, map[string]string{"s": s, "prefix": prefix}, strings.TrimPrefix(s, prefix))
	}
}

func TestTrimSuffix(t *testing.T) {
	suffix := "bar"
	for _, s := range []string{"foo", " foobar", "foo bar"} {
		executeTestTemplate(t, "title", `{{ trimSuffix(.s, .suffix) }}`, nil, map[string]string{"s": s, "suffix": suffix}, strings.TrimSuffix(s, suffix))
	}
}
