package portal

import "github.com/Bookshelf-Writer/esp32tool/esp32_old/code"

//###########################################################//

type MsgObj struct {
	Direction code.DirectionType
	Opcode    code.OpType

	Data   []byte
	Length int

	Checksum []byte
}
