package command

import (
	"bytes"
	"github.com/Bookshelf-Writer/esp32tool/esp32_old/code"
)

//###########################################################//

func Sync() *CommandObj {
	buf := initBuffer()

	buf.Write([]byte{0x07, 0x07, 0x12, 0x20})
	buf.Write(bytes.Repeat([]byte{0x55}, 32))

	return newRequest(code.OpSync, buf.Bytes())
}

func AttachSpiFlash() *CommandObj {
	return newRequest(code.OpSpiAttachFlash, make([]byte, 8))
}

//

func ChangeBaudRate(newBaudrate uint32, oldBaudrate uint32) *CommandObj {
	buf := initBuffer()

	buf.Uint32(newBaudrate)
	buf.Uint32(oldBaudrate)

	return newRequest(code.OpChangeBaudRate, buf.Bytes())
}