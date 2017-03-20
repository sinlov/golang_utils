package cli

import (
	"os/exec"
	"strings"
	"github.com/axgle/mahonia"
	"os"
)

// system use windows, other is linux or OSX
// return is success bool processState string and cmdOut string
func CmdExec(cmd string, system string) (bool, string, string) {
	var c *exec.Cmd
	var cmdOut string
	if system == "windows" {
		argArray := strings.Split("/c "+cmd, " ")
		c = exec.Command("cmd", argArray...)
	} else {
		c = exec.Command("/bin/sh", "-c", cmd)
	}
	out, _ := c.Output()
	cmdOut = string(out)
	if system == "windows" {
		dec := mahonia.NewDecoder("gbk")
		cmdOut = dec.ConvertString(cmdOut)
	}
	var processState os.ProcessState
	c.ProcessState = &processState
	processSuccess := processState.Success()
	processStateStr := processState.String()
	return processSuccess, processStateStr, cmdOut
}
