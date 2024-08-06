package esp32tool

import (
	"github.com/Bookshelf-Writer/esp32tool/esp32_old"
	"github.com/Bookshelf-Writer/esp32tool/esp32_old/command"
	"github.com/Bookshelf-Writer/esp32tool/lib/output"
	"github.com/Bookshelf-Writer/esp32tool/lib/serial"
	"time"
)

func ConnectEsp32(portPath string, connectBaudrate uint32, transferBaudrate uint32, retries uint, logger *output.LogObj) (*esp32_old.ESP32ROM, error) {
	logger = logger.NewLog("ConnectESP")

	serialPort, err := serial.NewEsp(portPath, int(connectBaudrate))
	if err != nil {
		logger.Trace().Err(err).Msg("serial.PortInit")
		return nil, err
	}

	esp := esp32_old.NewESP32ROM(serialPort, logger)
	err = esp.Connect(retries)
	if err != nil {
		logger.Trace().Err(err).Msg("esp32.Connect")
		return nil, err
	}

	//установка скорости подключения
	{
		_, err = esp32_old.CheckExecuteCommand(serialPort,
			command.ChangeBaudRate(transferBaudrate, 0),
			100*time.Millisecond,
			3,
		)
		if err != nil {
			return nil, err
		}

		err = serialPort.BaudRate.Set(int(transferBaudrate))
		if err != nil {
			return nil, err
		}

		time.Sleep(10 * time.Millisecond)

		err = serialPort.Flush() // get rid of crap sent during baud rate change
	}

	return esp, err
}
