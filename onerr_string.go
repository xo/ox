// Code generated by "stringer -type OnErr"; DO NOT EDIT.

package kobra

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OnErrExit-0]
	_ = x[OnErrContinue-1]
	_ = x[OnErrPanic-2]
}

const _OnErr_name = "OnErrExitOnErrContinueOnErrPanic"

var _OnErr_index = [...]uint8{0, 9, 22, 32}

func (i OnErr) String() string {
	if i >= OnErr(len(_OnErr_index)-1) {
		return "OnErr(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OnErr_name[_OnErr_index[i]:_OnErr_index[i+1]]
}
