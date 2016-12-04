package nohup

import (
	"os"
	"fmt"
	"os/exec"
)

/*
	If pid is child process will return nil
 */
func WatchPidFile(pidFilePath string) *exec.Cmd {
	pid := os.Getpid()
	fmt.Printf("this process pid %v\n", pid)
	if pid != 1 {
		cmd := exec.Command(pidFilePath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		return cmd
	}
	return nil
}
