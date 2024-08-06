package api

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NetworkInfo(ctx *gin.Context) {
	cmd := `ifconfig eth0 | grep "RX packets" | awk '{ print $5 }'`
	rx, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(string(rx))

	cmd = `ifconfig eth0 | grep "TX packets" | awk '{ print $5 }'`
	tx, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(string(tx))

	ctx.JSON(200, gin.H{"rx": string(rx), "tx": string(tx)})
}
