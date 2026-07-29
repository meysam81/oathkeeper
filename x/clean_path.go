// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"path"
	"strings"
)

// CleanPath resolves dot segments in p like path.Clean, but preserves a
// trailing slash so that the upstream receives the path as requested.
func CleanPath(p string) string {
	cleaned := path.Clean(p)
	if strings.HasSuffix(p, "/") && cleaned != "/" {
		cleaned += "/"
	}
	return cleaned
}
