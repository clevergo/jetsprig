// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jetsprig

import (
	"bytes"
	"testing"

	"github.com/CloudyKit/jet/v5"
	"github.com/stretchr/testify/assert"
)

var testSet *jet.Set

func TestMain(m *testing.M) {
	testSet = jet.NewHTMLSet("")
	GenericFuncMap().AttachTo(testSet)

	m.Run()
}

func executeTestTemplate(name, template string, vars jet.VarMap, data interface{}) (string, error) {
	t, err := testSet.Parse(name, template)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = t.Execute(&out, vars, data)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func runTest(t *testing.T, name, template string, vars jet.VarMap, data interface{}, expected string) {
	out, err := executeTestTemplate(name, template, vars, data)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, out)
}
