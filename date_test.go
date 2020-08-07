// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"bytes"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	tmpl, err := testSet.Parse("now", "{{ now().UnixNano() }}")
	if err != nil {
		t.Fatal(err)
	}

	startedAt := time.Now()
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	sec, err := strconv.ParseInt(buf.String(), 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, sec > startedAt.UnixNano())
	assert.True(t, sec < time.Now().UnixNano())
}
