package format

import "time"

func LayoutNowTime(layout string) string {
	return time.Now().Format(layout)
}
