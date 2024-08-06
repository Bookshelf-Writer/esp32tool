package esp32_old

import (
	"fmt"
	"github.com/Bookshelf-Writer/esp32tool/esp32_old/command"
	"github.com/Bookshelf-Writer/esp32tool/esp32_old/portal"
	"github.com/Bookshelf-Writer/esp32tool/lib/serial"
	"time"
)

//###########################################################//

func RunCommand(port *serial.SerialObj, command *command.CommandObj, timeout time.Duration) (*portal.ResponseObj, error) {
	err := portal.Write(port, command.Bytes())
	if err != nil {
		return nil, err
	}

	for retryCount := 0; retryCount < 16; retryCount++ {

		responseBuf, err := portal.Read(port, timeout)
		if err != nil {
			return nil, err
		}

		if responseBuf[1] != command.OpcodeToByte() {
			//e.log.Trace().Msg("Opcode did not match")
			continue
		} else {
			return portal.Response(responseBuf)
		}
	}

	return nil, fmt.Errorf("Retrycount exceeded")
}

func ReadRegister(port *serial.SerialObj, timeout time.Duration, register uint32) ([]byte, error) {
	response, err := RunCommand(port, command.Read.Register(register), timeout)
	if err != nil {
		return make([]byte, 4), err
	}

	return response.Checksum(), nil
}

func ReadEfuse(port *serial.SerialObj, timeout time.Duration, eFuseIndex uint32) ([]byte, error) {
	return ReadRegister(port, timeout, 0x6001a000+(4*eFuseIndex))
}

//////

func CheckExecuteCommand(port *serial.SerialObj, command *command.CommandObj, timeout time.Duration, retries int) (*portal.ResponseObj, error) {
	for retryCount := 0; retryCount < retries; retryCount++ {

		response, err := RunCommand(port, command, timeout)
		if err != nil {
			continue
		}

		if !response.Status {
			continue
		} else {
			return response, err
		}

	}

	return nil, fmt.Errorf("unexpected ending")
}