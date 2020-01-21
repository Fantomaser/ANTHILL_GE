package streams

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// MonitorRect is monitor real size
type MonitorRect struct {
	H int
	W int
}

// MonitorInfo struct of monitor info
type MonitorInfo struct {
	MonitorParam []string
	MonitorSize  []MonitorRect
}

// Monitor real struct of monitor info
var Monitor MonitorInfo

// ConfigurateRect Monitor configuration
func ConfigurateRect() {

	args := "Add-Type -AssemblyName System.Windows.Forms;[System.Windows.Forms.Screen]::AllScreens"
	out, err := exec.Command("powershell", args).Output()

	Monitor.MonitorParam = strings.SplitAfter(string(out), "\n")

	if err != nil {
		log.Println(err.Error())
	}

	for i := 0; i < (len(Monitor.MonitorParam)-3)/6; i++ {
		GetSize(i)
	}

}

// GetSize get size of monitor and make info structure
func GetSize(monitorNum int) {

	sizeH := make([]byte, 0)
	sizeW := make([]byte, 0)

	for i := 0; len(Monitor.MonitorParam[3+monitorNum*6]) > i; i++ {
		if Monitor.MonitorParam[3+monitorNum*6][i] == 'h' {
			for j := i + 2; Monitor.MonitorParam[3+monitorNum*6][j] != ','; j++ {
				sizeH = append(sizeH, Monitor.MonitorParam[3+monitorNum*6][j])
			}
			break
		}
	}

	for i := 0; len(Monitor.MonitorParam[3+monitorNum*6]) > i; i++ {
		if Monitor.MonitorParam[3+monitorNum*6][i] == 'g' {
			for j := i + 4; Monitor.MonitorParam[3+monitorNum*6][j] != '}'; j++ {
				sizeW = append(sizeW, Monitor.MonitorParam[3+monitorNum*6][j])
			}
			break
		}
	}

	h, _ := strconv.Atoi(string(sizeH))
	w, _ := strconv.Atoi(string(sizeW))

	Monitor.MonitorSize = append(Monitor.MonitorSize, MonitorRect{h, w})

	fmt.Println(h, w)
}
