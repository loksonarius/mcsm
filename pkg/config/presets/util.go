package presets

import (
	"fmt"
)

func e(s string, v ...interface{}) error {
	return fmt.Errorf(s, v...)
}
