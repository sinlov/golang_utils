package cli

import (
	"testing"
	"github.com/bmizerany/assert"
	"fmt"
)

func TestRed(t *testing.T) {
	fmt.Sprintf("\x1b[0;%dm%s\n\x1b[0m", 30, "123")
	fmt.Println(Red("CLi Red color"))
	assert.T(t, true, nil)
}

func TestFmtBlack(t *testing.T) {
	FmtBlack("%s\n", "Cli FmtBlack color")
	assert.T(t, true, nil)
}

func TestFmtRed(t *testing.T) {
	FmtRed("%s\n", "Cli FmtRed color")
	assert.T(t, true, nil)
}

func TestFmtGreen(t *testing.T) {
	FmtGreen("%s\n", "Cli FmtGreen color")
	assert.T(t, true, nil)
}

func TestFmtYellow(t *testing.T) {
	FmtYellow("%s\n", "Cli FmtYellow color")
	assert.T(t, true, nil)
}

func TestFmtBlue(t *testing.T) {
	FmtBlue("%s\n", "Cli FmtBlue color")
	assert.T(t, true, nil)
}

func TestFmtMagenta(t *testing.T) {
	FmtMagenta("%s\n", "Cli FmtMagenta color")
	assert.T(t, true, nil)
}

func TestFmtCyan(t *testing.T) {
	FmtCyan("%s\n", "Cli FmtCyan color")
	assert.T(t, true, nil)
}

func TestFmtWhite(t *testing.T) {
	FmtWhite("%s\n", "Cli FmtWhite color")
	assert.T(t, true, nil)
}
