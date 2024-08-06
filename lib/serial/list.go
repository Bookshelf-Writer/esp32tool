package serial

import (
	"go.bug.st/serial"
)

//###########################################################//

func ListSerial() ([]string, error) {
	return serial.GetPortsList()
}
