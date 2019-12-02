package cli

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCliColor(t *testing.T) {
	tests := []struct {
		name       string
		forTestStr string
		wantCliStr string
	}{
		{
			name:       "cli.Black",
			forTestStr: "CLi Black color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 30, "CLi Black color"),
		},
		{
			name:       "cli.Red",
			forTestStr: "CLi Red color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 31, "CLi Red color"),
		},
		{
			name:       "cli.Green",
			forTestStr: "CLi Green color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 32, "CLi Green color"),
		},
		{
			name:       "cli.Yellow",
			forTestStr: "CLi Yellow color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 33, "CLi Yellow color"),
		},
		{
			name:       "cli.Blue",
			forTestStr: "CLi Blue color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 34, "CLi Blue color"),
		},
		{
			name:       "cli.Magenta",
			forTestStr: "CLi Magenta color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 35, "CLi Magenta color"),
		},
		{
			name:       "cli.Cyan",
			forTestStr: "CLi Cyan color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 36, "CLi Cyan color"),
		},
		{
			name:       "cli.White",
			forTestStr: "CLi White color",
			wantCliStr: fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 37, "CLi White color"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.name {
			case "cli.Black":
				actResult := Black(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtBlack("%v%v", test.forTestStr, "\n")
			case "cli.Red":
				actResult := Red(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtRed("%v%v", test.forTestStr, "\n")
			case "cli.Green":
				actResult := Green(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtGreen("%v%v", test.forTestStr, "\n")
			case "cli.Yellow":
				actResult := Yellow(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtYellow("%v%v", test.forTestStr, "\n")
			case "cli.Blue":
				actResult := Blue(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtBlue("%v%v", test.forTestStr, "\n")
			case "cli.Magenta":
				actResult := Magenta(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtMagenta("%v%v", test.forTestStr, "\n")
			case "cli.Cyan":
				actResult := Cyan(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtCyan("%v%v", test.forTestStr, "\n")
			case "cli.White":
				actResult := White(test.forTestStr)
				assert.Equal(t, test.wantCliStr, actResult)
				FmtWhite("%v%v", test.forTestStr, "\n")
			}
		})
	}
}

//func TestFmtBlack(t *testing.T) {
//	convey.Convey("TestFmtBlack", t, func() {
//		FmtBlack("%s\n", "Cli FmtBlack color")
//		convey.So("Cli FmtBlack color", convey.ShouldEqual, "Cli FmtBlack color")
//	})
//}
//
//func TestFmtRed(t *testing.T) {
//	convey.Convey("TestFmtRed", t, func() {
//		FmtRed("%s\n", "Cli FmtRed color")
//		convey.So("Cli TestFmtRed color", convey.ShouldEqual, "Cli TestFmtRed color")
//	})
//}
//
//func TestFmtGreen(t *testing.T) {
//	convey.Convey("TestFmtGreen", t, func() {
//		FmtGreen("%s\n", "Cli FmtGreen color")
//		convey.So("Cli TestFmtGreen color", convey.ShouldEqual, "Cli TestFmtGreen color")
//	})
//}
//
//func TestFmtYellow(t *testing.T) {
//	convey.Convey("TestFmtYellow", t, func() {
//		FmtYellow("%s\n", "Cli FmtYellow color")
//		convey.So("Cli TestFmtYellow color", convey.ShouldEqual, "Cli TestFmtYellow color")
//	})
//}
//
//func TestFmtBlue(t *testing.T) {
//	convey.Convey("TestFmtBlue", t, func() {
//		FmtBlue("%s\n", "Cli FmtBlue color")
//		convey.So("Cli TestFmtBlue color", convey.ShouldEqual, "Cli TestFmtBlue color")
//	})
//}
//
//func TestFmtMagenta(t *testing.T) {
//	convey.Convey("TestFmtMagenta", t, func() {
//		FmtMagenta("%s\n", "Cli FmtMagenta color")
//		convey.So("Cli TestFmtMagenta color", convey.ShouldEqual, "Cli TestFmtMagenta color")
//	})
//}
//
//func TestFmtCyan(t *testing.T) {
//	convey.Convey("TestFmtCyan", t, func() {
//		FmtCyan("%s\n", "Cli FmtCyan color")
//		convey.So("Cli TestFmtCyan color", convey.ShouldEqual, "Cli TestFmtCyan color")
//	})
//}
//
//func TestFmtWhite(t *testing.T) {
//	convey.Convey("TestFmtWhite", t, func() {
//		FmtWhite("%s\n", "Cli FmtWhite color")
//		convey.So("Cli TestFmtWhite color", convey.ShouldEqual, "Cli TestFmtWhite color")
//	})
//}
