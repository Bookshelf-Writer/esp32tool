package esp32s2

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32 S2"

var MagicValues = []uint32{
	0x0000_07c6,
}

var (
	IROM = patterns.AddressObj{
		Offset: 0x4008_0000,
		Size:   0x40b8_0000 - 0x4008_0000,
	}
	DROM = patterns.AddressObj{
		Offset: 0x3f00_0000,
		Size:   0x3f3f_0000 - 0x3f00_0000,
	}
)

const Efuse uint32 = 0x3f41_a000
