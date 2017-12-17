package fastest

// True asserts that cond is true.
func (ft T) True(cond bool, msg ...interface{}) {
	if !cond {
		if len(msg) == 0 {
			ft.locatedFatalf("true != %v", cond)
		} else {
			ft.locatedFatalf(msg...)
		}
	}
}

//False asserts that cond is true.
func (ft T) False(cond bool, msg ...interface{}) {
	if cond {
		if len(msg) == 0 {
			ft.locatedFatalf("false != %v", cond)
		} else {
			ft.locatedFatalf(msg...)
		}
	}
}

// Equals asserts that a and b are equal.
func (ft T) Equals(a, b interface{}, msg ...interface{}) {
	if a != b {
		if len(msg) == 0 {
			ft.TypedFatalf("%v != %v", a, b)
		} else {
			ft.locatedFatalf(msg...)
		}
	}
}

// Implies implements a -> b.
// That means it will only fail when a is false and b is true.
func (ft T) Implies(a, b bool, msg ...interface{}) {
	if a && !b {
		if len(msg) == 0 {
			ft.locatedFatalf("%v !-> %v", a, b)
		} else {
			ft.locatedFatalf(msg...)
		}
	}
}

// Only skips the rest of the test when cond is false.
// All conditions that follow an Only() call can therefore assume that cond is
// true.
func (ft T) Only(cond bool) {
	if !cond {
		ft.SkipNow()
	}
}

// Nil asserts that parameter a is nil.
func (ft T) Nil(a interface{}, msg ...interface{}) {
	if a != nil {
		if len(msg) == 0 {
			ft.TypedFatalf("nil != %v", a)
		} else {
			ft.locatedFatalf(msg...)
		}
	}
}
