// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanPath(t *testing.T) {
	for _, tc := range []struct {
		in, out string
	}{
		{"/", "/"},
		{"/select/vmui", "/select/vmui"},
		{"/select/vmui/", "/select/vmui/"},
		{"/user/../admin/secrets", "/admin/secrets"},
		{"/user/../admin/secrets/", "/admin/secrets/"},
		{"/a/./b//c/", "/a/b/c/"},
		{"/..", "/"},
		{"/../", "/"},
	} {
		assert.Equal(t, tc.out, CleanPath(tc.in), "input: %s", tc.in)
	}
}
