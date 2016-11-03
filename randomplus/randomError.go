package randomplus

import (
	"fmt"
)

type RandomSizeError struct {
	code int
	msg  string
}

func (s RandomSizeError) Error() string {
	return fmt.Sprintf("randomplus: size error code %v, message %v", s.code, s.msg)
}
