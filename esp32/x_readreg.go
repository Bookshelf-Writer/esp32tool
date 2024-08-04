package esp32

import (
	"bytes"
	"github.com/Bookshelf-Writer/esp32tool/lib/serial"
	"time"
)

//###########################################################//

func GetUID(port *serial.SerialObj, timeout time.Duration) ([]byte, error) {
	var buf bytes.Buffer

	mac, err := ReadEfuse(port, timeout, 2)
	if err != nil {
		return nil, err
	}
	buf.Write(mac)

	mac, err = ReadEfuse(port, timeout, 1)
	if err != nil {
		return nil, err
	}
	buf.Write(mac)

	return buf.Bytes(), nil
}
