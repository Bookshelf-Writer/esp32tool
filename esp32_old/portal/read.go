package portal

import (
	"bytes"
	"github.com/Bookshelf-Writer/esp32tool/esp32_old/code"
	"github.com/Bookshelf-Writer/esp32tool/lib/serial"
	"time"
)

//###########################################################//

func Read(port *serial.SerialObj, timeout time.Duration) ([]byte, error) {
	state := code.StateWaitingHeader
	startTime := time.Now()
	var buf bytes.Buffer

	err := port.Timeout.Set(timeout)
	if err != nil {
		return nil, err
	}

	for {
		if time.Since(startTime) > timeout {
			return nil, ErrTimeoutRead
		}

		byteBuf := make([]byte, 1)
		n, err := port.Read(byteBuf)
		if err != nil {
			if err.Error() == "EOF" {
				continue
			}
			return nil, err
		}
		if n != 1 {
			continue
		}

		switch state {

		case code.StateWaitingHeader:
			if byteBuf[0] == code.SlipEnd.Byte() {
				state = code.StateReadingContent
			}

		case code.StateReadingContent:
			switch byteBuf[0] {
			case code.SlipEnd.Byte():
				return buf.Bytes(), nil

			case code.SlipEsc.Byte():
				state = code.StateInEscape
			default:
				buf.WriteByte(byteBuf[0])
			}

		case code.StateInEscape:
			switch byteBuf[0] {
			case 0xDC:
				buf.WriteByte(code.SlipEnd.Byte())
				state = code.StateReadingContent
			case 0xDD:
				buf.WriteByte(code.SlipEsc.Byte())
				state = code.StateReadingContent
			default:
				return nil, ErrUnexpectedChar
			}

		}
	}
}
