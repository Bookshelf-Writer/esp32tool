package code

/* This file is automatically generated */

type SlipType byte

func (obj SlipType) Byte() byte {
	return byte(obj)
}

const (
	SlipHeader     SlipType = 192 //0xc0
	SlipEscapeChar SlipType = 219 //0xdb
)

const (
	SlipTextHeader     = "Header"
	SlipTextEscapeChar = "Escape Char"
)

var SlipMap = map[SlipType]string{
	SlipHeader:     SlipTextHeader,
	SlipEscapeChar: SlipTextEscapeChar,
}

func (obj SlipType) String() string {
	val, ok := SlipMap[obj]
	if ok {
		return val
	}
	return "Unknown SlipType"
}