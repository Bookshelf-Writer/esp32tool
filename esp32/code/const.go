package code

//###########################################################//

const (
	DirectionRequest  byte = 0x00
	DirectionResponse byte = 0x01
)

const (
	SlipEnd    byte = 0xc0
	SlipEsc    byte = 0xdb
	SlipEscEnd byte = 0xdc
	SlipEscEsc byte = 0xdd
)

const (
	MagicRegAdr     uint32 = 0x40001000
	FlashSectorSize uint32 = 0x1000
	FlashWriteSize  uint32 = 0x400
)
