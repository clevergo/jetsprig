// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	startedAt := time.Now()
	out, err := executeTestTemplate("now", "{{ now().UnixNano() }}", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	sec, err := strconv.ParseInt(out, 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, sec > startedAt.UnixNano())
	assert.True(t, sec < time.Now().UnixNano())
}

func TestDate(t *testing.T) {
	layout := "2006-01-02"
	date := time.Now()
	runTest(t, "date", `{{ date(.date, .layout) }}`, nil, map[string]interface{}{"date": date, "layout": layout}, date.Format(layout))
}

func TestDateInZone(t *testing.T) {
	layout := "2006-01-02"
	date := time.Now()

	for _, test := range []struct {
		zone string
	}{
		{"Asia/Shanghai"},
		{"Asia/Nil"},
	} {
		loc, err := time.LoadLocation(test.zone)
		tmpl := `{{ dateInZone(.date, .zone) | date(_, .layout) }}`
		data := map[string]interface{}{
			"date":   date,
			"layout": layout,
			"zone":   test.zone,
		}
		if err != nil {
			_, err = executeTestTemplate("dateInZone", tmpl, nil, data)
			assert.NotNil(t, err)
		} else {
			runTest(t, "dateInZone", tmpl, nil, data, date.In(loc).Format(layout))
		}
	}
}

func TestAgo(t *testing.T) {
	date := time.Now().Add(-time.Second)
	out, err := executeTestTemplate("ago", `{{ ago(.date) }}`, nil, map[string]interface{}{"date": date})
	if err != nil {
		t.Fatal(err)
	}
	duration, err := time.ParseDuration(out)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, duration > time.Second)
	assert.True(t, duration < 2*time.Second)
}
