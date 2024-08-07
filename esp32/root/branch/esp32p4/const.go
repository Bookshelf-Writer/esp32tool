package esp32p4

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32 P4"

var MagicValues = []uint32{
	0x0,
}

var (
	IROM = patterns.AddressObj{
		Offset: 0x4000_0000,
		Size:   0x4C00_0000 - 0x4000_0000,
	}
	DROM = patterns.AddressObj{
		Offset: 0x4000_0000,
		Size:   0x4C00_0000 - 0x4000_0000,
	}
)

const Efuse uint32 = 0x5012_D000
