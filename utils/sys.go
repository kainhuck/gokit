package utils

import (
	"os"
	"os/exec"
)

func RunCmd(view bool, command string, args ...string) error{
	cmdP, err := exec.LookPath(command)
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdP, args...)
	if view {
		cmd.Stdout = os.Stdout
	}

	return cmd.Run()
}

func RunCmdOut(command string, args ...string)(output string,err error){
	cmdP, err := exec.LookPath(command)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(cmdP, args...)
	bts, err := cmd.Output()

	return string(bts), err
}
