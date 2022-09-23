package utilities

import (
	"fmt"
	"log"

	sh "github.com/abdfnx/shell"
	"github.com/itchyny/volume-go"
)

func ChangeBrightness(level string) {
	if level > 100 && level < 0 {
		log.Panic("Level should be more than 0 and less than 100")
	}

	toExec := fmt.Sprintf("(Get-WmiObject -Namespace root/WMI -Class WmiMonitorBrightnessMethods).WmiSetBrightness(1,%s)", level)

	sh.PWSLCmd(toExec)
}

func ChangeVolume(level int) {
	if level > 100 && level < 0 {
		log.Panic("Level should be more than 0 and less than 100")
	}

	if err := volume.SetVolume(level); err != nil {
		log.Panic(err)
	}
}

func MuteVolume() {
	if err := volume.Mute(); err != nil {
		log.Panic(err)
	}
}
