package fastest

import "testing"

// Seq runs a subtest sequentially.
// It is functionally equal to testing.T.Run().
func (ft T) Seq(s string, f func(ft T)) bool {
	return ft.Run(s, func(t *testing.T) {
		ft := T{T: t}
		f(ft)
	})
}

// Par runs a subtest in parallel.
// It is equivalent to
//   t.Run(s, func(t *testing.T){
//     t.Parallel()
//     // ...
//   })
func (ft T) Par(s string, f func(ft T)) bool {
	return ft.Run(s, func(t *testing.T) {
		ft := T{T: t}
		ft.Parallel()
		f(ft)
	})
}
