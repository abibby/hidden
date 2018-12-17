// +build !windows

package hidden

import "path/filepath"

func IsHidden(path string) (bool, error) {
	parts := filepath.SplitList(path)
	for _, part := range parts {
		if len(part) > 0 && part[0] == '.' {
			return true, nil
		}
	}
	return false, nil
}
