package code

//###########################################################//

type OpType byte

func (obj OpType) Byte() byte {
	return byte(obj)
}

const (
	OpFlashBegin OpType = 0x02
	OpFlashData  OpType = 0x03
	OpFlashEnd   OpType = 0x04
	OpMemBegin   OpType = 0x05
	OpMemEnd     OpType = 0x06
	OpMemData    OpType = 0x07
	OpSync       OpType = 0x08
	OpWriteReg   OpType = 0x09
	OpReadReg    OpType = 0x0A

	OpSpiSetParams   OpType = 0x0B
	OpSpiAttach      OpType = 0x0D
	OpChangeBaudrate OpType = 0x0F
	OpFlashDeflBegin OpType = 0x10
	OpFlashDeflData  OpType = 0x11
	OpFlashDeflEnd   OpType = 0x12
	OpFlashMd5       OpType = 0x13

	OpGetSecurityInfo OpType = 0x14

	OpEraseFlash  OpType = 0xD0
	OpEraseRegion OpType = 0xD1
	OpReadFlash   OpType = 0xD2
	OpRunUserCode OpType = 0xD3

	OpFlashEncryptedData OpType = 0xD4
	OpFlashDetect        OpType = 0x9F
)
