package fastest

//True asserts that cond is true.
func (ft T) True(cond bool, msg ...interface{}) {
	if !cond {
		if len(msg) == 0 {
			ft.locatedFatalf("true != %v", cond)
		} else {
			ft.locatedFatalf(msg)
		}
	}
}

//Nope asserts that cond is true.
func (ft T) Nope(cond bool, msg ...interface{}) {
	if cond {
		if len(msg) == 0 {
			ft.locatedFatalf("false != %v", cond)
		} else {
			ft.locatedFatalf(msg)
		}
	}
}

// Same asserts that a and b are equal.
func (ft T) Same(a interface{}, b interface{}, msg ...interface{}) {
	if a != b {
		if len(msg) == 0 {
			ft.TypedFatalf("%v != %v", a, b)
		} else {
			ft.locatedFatalf(msg)
		}
	}
}

// Cond implements a -> b.
// That means it will only fail when a is false and b is true.
func (ft T) Cond(a bool, b bool, msg ...interface{}) {
	if a && !b {
		if len(msg) == 0 {
			ft.locatedFatalf("%v !-> %v", a, b)
		} else {
			ft.locatedFatalf(msg)
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
			ft.TypedFatalf("nil != %v", a)
		} else {
			ft.locatedFatalf(msg)
		}
	}
}
