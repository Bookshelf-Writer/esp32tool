package portal

import (
	"bytes"
	"github.com/Bookshelf-Writer/esp32tool/esp32_old/code"
)

//###########################################################//

func encode(data []byte) []byte {
	var buf bytes.Buffer
	buf.WriteByte(code.SlipEnd.Byte())

	for _, b := range data {
		switch b {
		case code.SlipEsc.Byte():
			buf.Write([]byte{code.SlipEsc.Byte(), code.SlipEscEsc.Byte()})
		case code.SlipEnd.Byte():
			buf.Write([]byte{code.SlipEsc.Byte(), code.SlipEscEnd.Byte()})
		default:
			buf.WriteByte(b)
		}

	}

	buf.WriteByte(code.SlipEnd.Byte())
	return buf.Bytes()
}
