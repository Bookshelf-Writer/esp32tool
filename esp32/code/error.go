package code

import "errors"

//###########################################################//

type ErrType byte

const (
	ErrInvalidMessageReceived           ErrType = 0x05
	ErrBootloaderFailedToExecuteCommand ErrType = 0x06
	ErrReceivedMessageHasInvalidCRC     ErrType = 0x07

	ErrBootloaderFailedToWriteToFlash  ErrType = 0x08
	ErrBootloaderFailedToReadFromFlash ErrType = 0x09
	ErrInvalidLengthForFlashRead       ErrType = 0x0a
	ErrMalformedCompressedDataReceived ErrType = 0x0b

	ErrBadDataLength   ErrType = 0xc0
	ErrBadDataChecksum ErrType = 0xc1
	ErrBadBlockSize    ErrType = 0xc2
	ErrInvalidCommand  ErrType = 0xc3

	ErrSPIOperationFailed ErrType = 0xc4
	ErrSPIUnlockFailed    ErrType = 0xc5

	ErrUncompressData         ErrType = 0xc7
	ErrDidntReceiveEnoughData ErrType = 0xc8
	ErrReceivedTooMuchData    ErrType = 0xc9
	ErrOther                  ErrType = 0xff
)

var (
	ErrTextInvalidMessageReceived           = errors.New("Invalid message received")
	ErrTextBootloaderFailedToExecuteCommand = errors.New("Bootloader failed to execute command")
	ErrTextReceivedMessageHasInvalidCRC     = errors.New("Received message has invalid CRC")

	ErrTextBootloaderFailedToWriteToFlash  = errors.New("Bootloader failed to write to flash")
	ErrTextBootloaderFailedToReadFromFlash = errors.New("Bootloader failed to read from flash")
	ErrTextInvalidLengthForFlashRead       = errors.New("Invalid length for flash read")
	ErrTextMalformedCompressedDataReceived = errors.New("Malformed compressed data received")

	ErrTextBadDataLength   = errors.New("Bad data length")
	ErrTextBadDataChecksum = errors.New("Bad data checksum")
	ErrTextBadBlockSize    = errors.New("Bad block size")
	ErrTextInvalidCommand  = errors.New("Invalid command")

	ErrTextSPIOperationFailed = errors.New("SPI operation failed")
	ErrTextSPIUnlockFailed    = errors.New("SPI unlock failed")

	ErrTextUncompressData         = errors.New("Error when uncompressing the data")
	ErrTextDidntReceiveEnoughData = errors.New("Didn't receive enough data")
	ErrTextReceivedTooMuchData    = errors.New("Received too much data")
	ErrTextOther                  = errors.New("Other")
)

var ErrMap = map[ErrType]error{
	ErrInvalidMessageReceived:           ErrTextInvalidMessageReceived,
	ErrBootloaderFailedToExecuteCommand: ErrTextBootloaderFailedToExecuteCommand,
	ErrReceivedMessageHasInvalidCRC:     ErrTextReceivedMessageHasInvalidCRC,
	ErrBootloaderFailedToWriteToFlash:   ErrTextBootloaderFailedToWriteToFlash,
	ErrBootloaderFailedToReadFromFlash:  ErrTextBootloaderFailedToReadFromFlash,
	ErrInvalidLengthForFlashRead:        ErrTextInvalidLengthForFlashRead,
	ErrMalformedCompressedDataReceived:  ErrTextMalformedCompressedDataReceived,
	ErrBadDataLength:                    ErrTextBadDataLength,
	ErrBadDataChecksum:                  ErrTextBadDataChecksum,
	ErrBadBlockSize:                     ErrTextBadBlockSize,
	ErrInvalidCommand:                   ErrTextInvalidCommand,
	ErrSPIOperationFailed:               ErrTextSPIOperationFailed,
	ErrSPIUnlockFailed:                  ErrTextSPIUnlockFailed,
	ErrUncompressData:                   ErrTextUncompressData,
	ErrDidntReceiveEnoughData:           ErrTextDidntReceiveEnoughData,
	ErrReceivedTooMuchData:              ErrTextReceivedTooMuchData,
	ErrOther:                            ErrTextOther,
}

func ErrorDetect(b byte) ErrType {
	_, ok := ErrMap[ErrType(b)]
	if ok {
		return ErrType(b)
	}
	return ErrOther
}
