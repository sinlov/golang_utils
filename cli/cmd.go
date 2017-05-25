package cli

import (
	"os/exec"
	"strings"
	"os"
	"runtime"
	"github.com/axgle/mahonia"
)

func getSystem() string {
	return runtime.GOOS
}

func IsSysWindows() bool {
	return getSystem() == "windows"
}

// chartSet "" default in windows "gbk", other is "utf-8"
// return isSuccess bool PID processState string and cmdOut string
func CmdExec(chartSet string, cmd ...string) (bool, int, string, string) {
	var c *exec.Cmd
	cmdStr := make([]string, 8)
	for i, ct := range cmd {
		trim := strings.Trim(ct, " ")
		cmdStr[i] = trim
	}
	cmdString := strings.Join(cmdStr, " ")
	var cmdOut string
	if IsSysWindows() {
		argArray := strings.Split("/c "+cmdString, " ")
		c = exec.Command("cmd", argArray...)
	} else {
		c = exec.Command("/bin/sh", "-c", cmdString)
	}
	//c.Stderr = os.Stderr do not set
	out, err := c.CombinedOutput()
	var processState os.ProcessState
	c.ProcessState = &processState
	processSuccess := processState.Success()
	processPid := processState.Pid()
	processStateStr := processState.String()
	cmdOut = string(out)
	if IsSysWindows() {
		if chartSet == "" {
			chartSet = "gbk"
		}
		dec := mahonia.NewDecoder(chartSet)
		cmdOut = dec.ConvertString(cmdOut)
	}
	if err != nil {
		return false, processPid, err.Error(), cmdOut
	}
	return processSuccess, processPid, processStateStr, cmdOut
}
