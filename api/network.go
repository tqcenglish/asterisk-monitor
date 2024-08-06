package api

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"

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

	rxInt, _ := strconv.Atoi(strings.TrimSuffix(string(rx), "\n"))
	txInt, _ := strconv.Atoi(strings.TrimSuffix(string(tx), "\n"))
	ctx.JSON(200, gin.H{"rx": rxInt / 1024 / 1024 / 1024, "tx": txInt / 1024 / 1024 / 1024})
}

func prettyByteSize(b int) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}
