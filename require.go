package fastest

//True asserts that cond is true.
func (ft T) True(cond bool, msg ...interface{}) {
	if !cond {
		if len(msg) == 1 {
			ft.Error(msg)
		} else {
			ft.Error("")
		}
	}
}

// Equal asserts that a and b are equal.
func (ft T) Equal(a interface{}, b interface{}, msg ...interface{}) {
	if a != b {
		ft.Error(msg)
	}
}

// Cond implements a -> b.
// That means it will only fail when a is false and b is true.
func (ft T) Cond(a bool, b bool, msg ...interface{}) {
	if a && !b {
		ft.Error(msg)
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
