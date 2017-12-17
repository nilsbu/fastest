package fastest

//True asserts that cond is true.
func (ft T) True(cond bool, msg ...interface{}) {
	if !cond {
		if len(msg) == 0 {
			ft.locatedErrorf("true != %v", cond)
		} else {
			ft.locatedErrorf(msg)
		}
	}
}

//Nope asserts that cond is true.
func (ft T) Nope(cond bool, msg ...interface{}) {
	if cond {
		if len(msg) == 0 {
			ft.locatedErrorf("false != %v", cond)
		} else {
			ft.locatedErrorf(msg)
		}
	}
}

// Same asserts that a and b are equal.
func (ft T) Same(a interface{}, b interface{}, msg ...interface{}) {
	if a != b {
		if len(msg) == 0 {
			ft.TypedErrorf("%v != %v", a, b)
		} else {
			ft.locatedErrorf(msg)
		}
	}
}

// Cond implements a -> b.
// That means it will only fail when a is false and b is true.
func (ft T) Cond(a bool, b bool, msg ...interface{}) {
	if a && !b {
		if len(msg) == 0 {
			ft.locatedErrorf("%v !-> %v", a, b)
		} else {
			ft.locatedErrorf(msg)
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

// Null ..
func (ft T) Null(a interface{}, msg ...interface{}) {
	if a != nil {
		if len(msg) == 0 {
			ft.TypedErrorf("nil != %v", a)
		} else {
			ft.locatedErrorf(msg)
		}
	}
}
