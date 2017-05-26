package cli

import (
	"os/exec"
	"strings"
	"os"
	"runtime"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"fmt"
	"github.com/pkg/errors"
)

var exitWithOutZero = errors.New("Has not run command exit code with 0")

type CmdTea struct {
	ChartSet   string
	IsPrint    bool
	CmdStrings string

	ErrorInfo error
	IsSuccess bool
	Pid       int
	ExitState string
	Out       string
	Err       string

	Env       []string
	ShellPath string
	Dir       string
	Args      []string
}

func getSystem() string {
	return runtime.GOOS
}

func IsSysWindows() bool {
	return getSystem() == "windows"
}

// chartSet "" default in windows "gbk", other is "utf-8"
// isPrint default is false
// return isSuccess bool PID processState string and cmdOut string
func (ct *CmdTea) CmdTeaInit(chartSet string, isPrint bool, cmd ...string) {
	cmdStr := make([]string, 8)
	for i, ct := range cmd {
		trim := strings.Trim(ct, " ")
		cmdStr[i] = trim
	}
	cmdString := strings.Join(cmdStr, " ")
	ct.ChartSet = chartSet
	ct.IsPrint = isPrint
	ct.CmdStrings = cmdString
}

func (ct CmdTea) CmdTeaRun() (bool, CmdTea) {
	if ct.CmdStrings == "" {
		ct.ErrorInfo = errors.New("You Cmd is Empty!")
		return false, ct

	}
	var c *exec.Cmd
	if IsSysWindows() {
		argArray := strings.Split("/c "+ct.CmdStrings, " ")
		c = exec.Command("cmd", argArray...)
	} else {
		c = exec.Command("/bin/sh", "-c", ct.CmdStrings)
	}
	out, combinedErr := c.CombinedOutput()

	ct.Env = c.Env
	ct.ShellPath = c.Path
	ct.Dir = c.Dir
	ct.Args = c.Args

	var processState os.ProcessState
	c.ProcessState = &processState
	processPid := processState.Pid()
	ct.Pid = processPid
	var dec mahonia.Decoder
	if ct.ChartSet == "" {
		if IsSysWindows() {
			ct.ChartSet = "gbk"
		} else {
			ct.ChartSet = "utf-8"
		}
	}
	dec = mahonia.NewDecoder(ct.ChartSet)

	cmdInterfaceOut := dec.ConvertString(string(out))
	if ct.IsPrint {
		fmt.Println(cmdInterfaceOut)

	}
	if combinedErr != nil {
		ct.IsSuccess = false
		ct.Err = cmdInterfaceOut
		ct.ErrorInfo = exitWithOutZero
		ct.ExitState = combinedErr.Error()
		return false, ct
	} else {
		ct.IsSuccess = true
		ct.Out = cmdInterfaceOut
		ct.ExitState = processState.String()
		return true, ct
	}
}

func CmdExec(chartSet string, cmd ...string) (bool, int, string, string) {
	var c *exec.Cmd
	cmdStr := make([]string, 8)
	for i, ct := range cmd {
		trim := strings.Trim(ct, " ")
		cmdStr[i] = trim
	}
	cmdString := strings.Join(cmdStr, " ")

	if IsSysWindows() {
		argArray := strings.Split("/c "+cmdString, " ")
		c = exec.Command("cmd", argArray...)
	} else {
		c = exec.Command("/bin/sh", "-c", cmdString)
	}

	var cmdOut string
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

func CmdRun(chartSet string, cmd ...string) (bool, error) {
	var c *exec.Cmd
	cmdStr := make([]string, 8)
	for i, ct := range cmd {
		trim := strings.Trim(ct, " ")
		cmdStr[i] = trim
	}
	cmdString := strings.Join(cmdStr, " ")

	if IsSysWindows() {
		argArray := strings.Split("/c "+cmdString, " ")
		c = exec.Command("cmd", argArray...)
	} else {
		c = exec.Command("/bin/sh", "-c", cmdString)
	}
	if chartSet == "" {
		if IsSysWindows() {
			chartSet = "gbk"
		} else {
			chartSet = "utf-8"
		}
	}
	dec := mahonia.NewDecoder(chartSet)
	stdout, err := c.StdoutPipe() //指向cmd命令的stdout
	stdErr, stderrErr := c.StderrPipe()
	c.Start()
	content, err := ioutil.ReadAll(stdout)
	contentErr, stderrErr := ioutil.ReadAll(stdErr)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if stderrErr != nil {
		fmt.Println(stderrErr)
		return false, stderrErr
	}
	if len(contentErr) > 0 {
		fmt.Println(Red(dec.ConvertString(string(contentErr))))
		return false, exitWithOutZero
	} else {
		fmt.Println(dec.ConvertString(string(content)))
		return true, nil
	}
}
