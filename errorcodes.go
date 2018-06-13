package fastest

// Code is and error code.
type Code int

const (
	// OK is the error code for no error.
	OK Code = iota
	// Fail is the error code for failure.
	Fail
	// Quit is the error code for abortion.
	Quit
)
