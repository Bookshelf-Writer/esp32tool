package code

/* This file is automatically generated */

type ErrType byte

const (
	ErrReceivedMessageIsInvalid     ErrType = 5 //0x05
	ErrFailedToActOnReceivedMessage ErrType = 6 //0x06
	ErrInvalidCRCInMessage          ErrType = 7 //0x07

	ErrFlashWriteError      ErrType = 8  //0x08
	ErrFlashReadError       ErrType = 9  //0x09
	ErrFlashReadLengthError ErrType = 10 //0x0a

	ErrDeflateError ErrType = 11 //0x0b
)

const (
	ErrTextReceivedMessageIsInvalid     = "Received message is invalid"
	ErrTextFailedToActOnReceivedMessage = "Failed to act on received message"
	ErrTextInvalidCRCInMessage          = "Invalid CRC in message"

	ErrTextFlashWriteError      = "Flash write error"
	ErrTextFlashReadError       = "Flash read error"
	ErrTextFlashReadLengthError = "Flash read length error"

	ErrTextDeflateError = "Deflate error"
)

var ErrMap = map[ErrType]string{
	ErrReceivedMessageIsInvalid:     ErrTextReceivedMessageIsInvalid,
	ErrFailedToActOnReceivedMessage: ErrTextFailedToActOnReceivedMessage,
	ErrInvalidCRCInMessage:          ErrTextInvalidCRCInMessage,
	ErrFlashWriteError:              ErrTextFlashWriteError,
	ErrFlashReadError:               ErrTextFlashReadError,
	ErrFlashReadLengthError:         ErrTextFlashReadLengthError,
	ErrDeflateError:                 ErrTextDeflateError,
}

func (obj ErrType) String() string {
	val, ok := ErrMap[obj]
	if ok {
		return val
	}
	return "Unknown ErrType"
}