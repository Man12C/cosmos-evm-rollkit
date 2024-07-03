// Code generated by "stringer -type=ConfLevel -linecomment"; DO NOT EDIT.

package xchain

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ConfUnknown-0]
	_ = x[ConfLatest-1]
	_ = x[ConfFast-2]
	_ = x[ConfSafe-3]
	_ = x[ConfFinalized-4]
	_ = x[confSentinel-5]
}

const _ConfLevel_name = "unknownlatestfastsafefinalsentinel must always be last"

var _ConfLevel_index = [...]uint8{0, 7, 13, 17, 21, 26, 54}

func (i ConfLevel) String() string {
	if i >= ConfLevel(len(_ConfLevel_index)-1) {
		return "ConfLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ConfLevel_name[_ConfLevel_index[i]:_ConfLevel_index[i+1]]
}
