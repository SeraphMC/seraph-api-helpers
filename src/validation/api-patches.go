package validation

import (
	"github.com/r3labs/diff/v3"
)

// DisplayPatches generates and applies patches to reflect the differences between the original and updated values.
func DisplayPatches[T any](original T, updated T) (diff.PatchLog, error) {
	if changelog, err := diff.Diff(original, updated, diff.DisableStructValues(), diff.AllowTypeMismatch(true)); err != nil {
		return nil, err
	} else {
		patches := new(T)
		patchLog := diff.Patch(changelog, patches)
		return patchLog, nil
	}
}
