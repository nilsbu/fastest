package fastest

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// TypedErrorf fails the test like t.Errorf but the message formats values in
// the format "type(Value)".
func (ft T) TypedErrorf(pattern string, vars ...interface{}) {

	s := sprintft(pattern, vars...)
	ft.locatedError(s)
}

func (ft T) locatedErrorf(msg ...interface{}) {
	ft.Errorf("%s: %s", getErrorLocation(),
		fmt.Sprintf(msg[0].(string), msg[1:]...))
}

func (ft T) locatedError(s string) {
	ft.Errorf("%s: %s", getErrorLocation(), s)
}

func sprintft(pattern string, vars ...interface{}) string {
	typed := make([]interface{}, len(vars))
	for i, v := range vars {
		typed[i] = fmt.Sprintf("%v(%v)", reflect.TypeOf(v), v)
	}

	s := fmt.Sprintf(pattern, typed...)
	return s
}

func getErrorLocation() string {
	for i := 2; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil || f.Name() == "testing.tRunner" {
			break
		}

		fileNameParts := strings.Split(file, "/")
		fileName := fileNameParts[len(fileNameParts)-1]

		if len(fileNameParts) >= 3 {
			packageName := fileNameParts[len(fileNameParts)-2]
			creatorName := fileNameParts[len(fileNameParts)-3]
			if creatorName == "nilsbu" && packageName == "fastest" {
				continue
			}
		}

		return fmt.Sprintf("%v:%v", fileName, line)
	}

	return "unknown origin"
}
