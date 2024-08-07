package esp32

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32"

var MagicValues = []uint32{
	0x00f0_1d83,
}

var (
	IROM = patterns.AddressObj{
		Offset: 0x400d_0000,
		Size:   0x4040_0000 - 0x400d_0000,
	}
	DROM = patterns.AddressObj{
		Offset: 0x3f40_0000,
		Size:   0x3f80_0000 - 0x3f40_0000,
	}
)

const Efuse uint32 = 0x3ff5_a000
