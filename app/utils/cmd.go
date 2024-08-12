package utils

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

var TimeZone string

// ExecCmdAsync 执行指定命令
func ExecCmdAsync(cmdName string, arg ...string) (stdOut, errOut string, err error) {
	logrus.Tracef("ExecCmdAsync cmd %s %s", cmdName, arg)
	cmd := exec.Command(cmdName, arg...)
	cmd.Env = append(os.Environ(), "TZ="+TimeZone)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		logrus.Errorf("cmd.Run(%s) failed with %s\n", cmdName, err)
	}
	outStr, errStr := stdout.String(), stderr.String()
	if len(outStr) > 0 {
		logrus.Debugf("cmd.Run(%s) %s", cmdName, outStr)
	}
	if len(errStr) > 0 {
		logrus.Errorf("cmd.Run(%s)%s", cmdName, errStr)
	}
	return outStr, errStr, err
}

// ExecCmd 执行指定命令
func ExecCmd(cmdName string, arg ...string) (stdOut, errOut string, err error) {
	logrus.Tracef("ExecCmd cmd %s %s", cmdName, arg)
	cmd := exec.Command(cmdName, arg...)
	cmd.Env = append(os.Environ(), "TZ="+TimeZone)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = cmd.Start()
	if err != nil {
		logrus.Error(err)
		return "", "", err
	}

	go func() {
		err = cmd.Wait()
		if err != nil {
			logrus.Errorf("cmd.Wait(%s) %s\n", cmdName, err)
		} else {
			logrus.Tracef("cmd.Wait(%s)\n", cmdName)
		}
		cancel()
	}()

	<-ctx.Done()

	outStr, errStr := stdout.String(), stderr.String()
	if len(outStr) > 0 {
		logrus.Debugf("cmd.Start(%s) %s", cmdName, outStr)
	}
	if len(errStr) > 0 {
		logrus.Errorf("cmd.Start(%s)%s", cmdName, errStr)
	}
	return outStr, errStr, err
}

func Reboot(sleep time.Duration) {
	go func() {
		logrus.Infof("start reboot and sleep: %d s", sleep)
		time.Sleep(sleep * time.Second)
		cmd := "/sbin/reboot"
		output, _ := exec.Command("bash", "-c", cmd).CombinedOutput()
		logrus.Info(cmd, string(output))
	}()
}

func RebootWeb(sleep time.Duration) {
	go func() {
		logrus.Infof("start rebootweb and sleep: %d s", sleep)
		time.Sleep(sleep * time.Second)
		output, err := exec.Command("/usr/local/bin/pmon2", "restart", "asterisk-gateway").CombinedOutput()
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info("pmon2 restart web", string(output))
	}()
}

func RunCmd(cmd string) (output []byte, err error) {
	logrus.Debug(cmd)
	output, err = exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		logrus.Error(err)
	}
	return
}
