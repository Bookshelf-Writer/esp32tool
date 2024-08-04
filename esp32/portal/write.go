package portal

import (
	"github.com/Bookshelf-Writer/esp32tool/lib/serial"
)

//###########################################################//

func Write(port *serial.SerialObj, data []byte) error {
	data = encode(data)

	n, err := port.Write(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return ErrMismatchBytes
	}

	return nil
}
