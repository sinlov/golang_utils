package cli

import (
	"testing"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
)

func TestRed(t *testing.T) {
	convey.Convey("TestRed", t, func() {
		// mock
		testString := "CLi Red color"
		redString := fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", 31, testString)
		convey.Convey("do TestRed", func() {
			// do
			redCliString := Red(testString)
			convey.Convey("verify TestRed", func() {
				// verify
				convey.So(redString, convey.ShouldEqual, redCliString)
			})
		})
	})
}

func TestFmtBlack(t *testing.T) {
	convey.Convey("TestFmtBlack", t, func() {
		FmtBlack("%s\n", "Cli FmtBlack color")
		convey.So("Cli FmtBlack color", convey.ShouldEqual, "Cli FmtBlack color")
	})
}

func TestFmtRed(t *testing.T) {
	convey.Convey("TestFmtRed", t, func() {
		FmtRed("%s\n", "Cli FmtRed color")
		convey.So("Cli TestFmtRed color", convey.ShouldEqual, "Cli TestFmtRed color")
	})
}

func TestFmtGreen(t *testing.T) {
	convey.Convey("TestFmtGreen", t, func() {
		FmtGreen("%s\n", "Cli FmtGreen color")
		convey.So("Cli TestFmtGreen color", convey.ShouldEqual, "Cli TestFmtGreen color")
	})
}

func TestFmtYellow(t *testing.T) {
	convey.Convey("TestFmtYellow", t, func() {
		FmtYellow("%s\n", "Cli FmtYellow color")
		convey.So("Cli TestFmtYellow color", convey.ShouldEqual, "Cli TestFmtYellow color")
	})
}

func TestFmtBlue(t *testing.T) {
	convey.Convey("TestFmtBlue", t, func() {
		FmtBlue("%s\n", "Cli FmtBlue color")
		convey.So("Cli TestFmtBlue color", convey.ShouldEqual, "Cli TestFmtBlue color")
	})
}

func TestFmtMagenta(t *testing.T) {
	convey.Convey("TestFmtMagenta", t, func() {
		FmtMagenta("%s\n", "Cli FmtMagenta color")
		convey.So("Cli TestFmtMagenta color", convey.ShouldEqual, "Cli TestFmtMagenta color")
	})
}

func TestFmtCyan(t *testing.T) {
	convey.Convey("TestFmtCyan", t, func() {
		FmtCyan("%s\n", "Cli FmtCyan color")
		convey.So("Cli TestFmtCyan color", convey.ShouldEqual, "Cli TestFmtCyan color")
	})
}

func TestFmtWhite(t *testing.T) {
	convey.Convey("TestFmtWhite", t, func() {
		FmtWhite("%s\n", "Cli FmtWhite color")
		convey.So("Cli TestFmtWhite color", convey.ShouldEqual, "Cli TestFmtWhite color")
	})
}
