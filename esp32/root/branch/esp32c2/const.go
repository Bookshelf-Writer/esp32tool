package esp32c2

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32 C2"

var MagicValues = []uint32{
	0x6f51_306f, // ECO0
	0x7c41_a06f, // ECO1
}

var (
	IROM = patterns.AddressObj{
		Offset: 0x4200_0000,
		Size:   0x4240_0000 - 0x4200_0000,
	}
	DROM = patterns.AddressObj{
		Offset: 0x3c00_0000,
		Size:   0x3c40_0000 - 0x3c00_0000,
	}
)

const Efuse uint32 = 0x6000_8800
