package cli

import (
	"testing"
	"github.com/bmizerany/assert"
	"fmt"
)

func TestBlack(t *testing.T) {
	fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 30, "123")
	fmt.Println(Red("CLi Red color"))
	assert.T(t, true, nil)
}
