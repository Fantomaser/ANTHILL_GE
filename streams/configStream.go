package streams

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type MonitorRect struct {
	H int
	W int
}

type MonitorInfo struct {
	MonitorParam []string
	MonitorSize  []MonitorRect
}

var Monitor MonitorInfo

//ConfigurateRect Monitor configuration
func ConfigurateRect() {
	args := "Add-Type -AssemblyName System.Windows.Forms;[System.Windows.Forms.Screen]::AllScreens"
	out, err := exec.Command("powershell", args).Output()

	Monitor.MonitorParam = strings.SplitAfter(string(out), "\n")

	if err != nil {
		log.Fatalln(err)
	}

	GetSize()

}

func GetSize() {

	sizeH := make([]byte, 0)
	sizeW := make([]byte, 0)

	for i := 0; len(Monitor.MonitorParam[3]) > i; i++ {
		if Monitor.MonitorParam[3][i] == 'h' {
			for j := i + 2; Monitor.MonitorParam[3][j] != ','; j++ {
				sizeH = append(sizeH, Monitor.MonitorParam[3][j])
			}
			break
		}
	}

	for i := 0; len(Monitor.MonitorParam[3]) > i; i++ {
		if Monitor.MonitorParam[3][i] == 'g' {
			for j := i + 4; Monitor.MonitorParam[3][j] != '}'; j++ {
				sizeW = append(sizeW, Monitor.MonitorParam[3][j])
			}
			break
		}
	}

	h, _ := strconv.Atoi(string(sizeH))
	w, _ := strconv.Atoi(string(sizeW))

	Monitor.MonitorSize = append(Monitor.MonitorSize, MonitorRect{h, w})

	fmt.Println(h, w)
}
