package cli

import (
	"runtime"
	"fmt"
)

const (
	TextBlack   = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}

func textColor(color int, str string) string {
	if IsWindows() {
		return str
	}

	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

func Black(str string) string {
	return textColor(TextBlack, str)
}

func FmtBlack(str string, a ... interface{}) {
	fmt.Print(Black(fmt.Sprintf(str, a)))
}

func Red(str string) string {
	return textColor(TextRed, str)
}

func FmtRed(str string, a ... interface{}) {
	fmt.Print(Red(fmt.Sprintf(str, a)))
}

func Green(str string) string {
	return textColor(TextGreen, str)
}

func FmtGreen(str string, a ... interface{}) {
	fmt.Print(Green(fmt.Sprintf(str, a)))
}

func Yellow(str string) string {
	return textColor(TextYellow, str)
}

func FmtYellow(str string, a ... interface{}) {
	fmt.Print(Yellow(fmt.Sprintf(str, a)))
}

func Blue(str string) string {
	return textColor(TextBlue, str)
}

func FmtBlue(str string, a ... interface{}) {
	fmt.Print(Blue(fmt.Sprintf(str, a)))
}

func Magenta(str string) string {
	return textColor(TextMagenta, str)
}

func FmtMagenta(str string, a ... interface{}) {
	fmt.Print(Magenta(fmt.Sprintf(str, a)))
}

func Cyan(str string) string {
	return textColor(TextCyan, str)
}

func FmtCyan(str string, a ... interface{}) {
	fmt.Print(Cyan(fmt.Sprintf(str, a)))
}

func White(str string) string {
	return textColor(TextWhite, str)
}

func FmtWhite(str string, a ... interface{}) {
	fmt.Print(White(fmt.Sprintf(str, a)))
}
