package code

//###########################################################//

type SlipType byte

func (obj SlipType) Byte() byte {
	return byte(obj)
}

const (
	SlipEnd    SlipType = 192 //0xc0
	SlipEsc    SlipType = 219 //0xdb
	SlipEscEnd SlipType = 220 //0xdc
	SlipEscEsc SlipType = 221 //0xdd
)
