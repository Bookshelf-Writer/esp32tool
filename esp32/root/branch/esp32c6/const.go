package esp32c6

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32 C6"

var MagicValues = []uint32{
	0x2CE0_806F,
}

var (
	IROM = patterns.AddressObj{
		Offset: 0x4200_0000,
		Size:   0x4280_0000 - 0x4200_0000,
	}
	DROM = patterns.AddressObj{
		Offset: 0x4280_0000,
		Size:   0x4300_0000 - 0x4280_0000,
	}
)

const Efuse uint32 = 0x600B_0800
