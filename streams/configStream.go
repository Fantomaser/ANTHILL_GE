package streams

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// ConfigurateRect Monitor configuration
func ConfigurateRect(monitor *MonitorInfo) {

	switch runtime.GOOS {
	case "windows":
		args := "Add-Type -AssemblyName System.Windows.Forms;[System.Windows.Forms.Screen]::AllScreens"
		out, err := exec.Command("powershell", args).Output()
		monitor.MonitorParam = strings.SplitAfter(string(out), "\n")
		if err != nil {
			log.Println(err.Error())
		}
		for i := 0; i < (len(monitor.MonitorParam)-3)/6; i++ {
			GetWinWndSize(i, monitor)
		}
	default:
		panic("Can not find OS")
	}

}

// GetWinWndSize get size of monitor and make info structure
func GetWinWndSize(monitorNum int, monitor *MonitorInfo) {

	sizeH := make([]byte, 0)
	sizeW := make([]byte, 0)

	for i := 0; len(monitor.MonitorParam[3+monitorNum*6]) > i; i++ {
		if monitor.MonitorParam[3+monitorNum*6][i] == 'h' {
			for j := i + 2; monitor.MonitorParam[3+monitorNum*6][j] != ','; j++ {
				sizeW = append(sizeW, monitor.MonitorParam[3+monitorNum*6][j])
			}
			break
		}
	}

	for i := 0; len(monitor.MonitorParam[3+monitorNum*6]) > i; i++ {
		if monitor.MonitorParam[3+monitorNum*6][i] == 'g' {
			for j := i + 4; monitor.MonitorParam[3+monitorNum*6][j] != '}'; j++ {
				sizeH = append(sizeH, monitor.MonitorParam[3+monitorNum*6][j])
			}
			break
		}
	}

	h, _ := strconv.Atoi(string(sizeH))
	w, _ := strconv.Atoi(string(sizeW))

	monitor.MonitorSize = append(monitor.MonitorSize, MonitorRect{h, w})

	fmt.Println("H: ",h,"\n W :", w)
}
