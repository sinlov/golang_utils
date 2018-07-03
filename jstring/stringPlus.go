package jstring

import "strings"

// if string s not empty
// check fist of string
// start is empty will return true
func StringStartWith(s string, start string) bool {
	s = strings.TrimSpace(s)
	start = strings.TrimSpace(start)
	sLen := len(s)
	checkLen := len(start)
	if checkLen == 0 {
		return true
	}
	if sLen == 0 {
		return false
	} else {
		if checkLen > sLen {
			return false
		} else {
			return s[0:checkLen] == start
		}
	}
}

// if string s not empty
// check end of string
// end is empty will return true
func StringEndWith(s string, end string) bool {
	s = strings.TrimSpace(s)
	end = strings.TrimSpace(end)
	sLen := len(s)
	checkLen := len(end)
	if checkLen == 0 {
		return true
	}
	if sLen == 0 {
		return false
	} else {
		if checkLen > sLen {
			return false
		} else {
			return s[sLen-checkLen:sLen] == end
		}
	}
}
