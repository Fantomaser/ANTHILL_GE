package streams

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var monitorInfo string
var H, W int

//ConfigurateRect Monitor configuration
func ConfigurateRect() {
	args := "Add-Type -AssemblyName System.Windows.Forms;[System.Windows.Forms.Screen]::AllScreens"
	out, err := exec.Command("powershell", args).Output()

	monitorInfo = strings.SplitAfter(string(out), "\n")[3]

	if err != nil {
		log.Fatalln(err)
	}

	GetSize()

}

func GetSize() {

	sizeH := make([]byte, 0)
	sizeW := make([]byte, 0)

	for i := 0; len(monitorInfo) > i; i++ {
		if monitorInfo[i] == 'h' {
			for j := i + 2; monitorInfo[j] != ','; j++ {
				sizeH = append(sizeH, monitorInfo[j])
			}
			break
		}
	}

	for i := 0; len(monitorInfo) > i; i++ {
		if monitorInfo[i] == 'g' {
			for j := i + 4; monitorInfo[j] != '}'; j++ {
				sizeW = append(sizeW, monitorInfo[j])
			}
			break
		}
	}

	H, _ = strconv.Atoi(string(sizeH))
	W, _ = strconv.Atoi(string(sizeW))

	fmt.Println(H, W)
}
