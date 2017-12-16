package fastest

import "testing"

// T is the test object.
// It provides full downward compatability with *testing.T.
type T struct {
	*testing.T
}
