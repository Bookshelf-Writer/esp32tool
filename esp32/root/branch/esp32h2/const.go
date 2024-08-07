package esp32h2

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32 H2"

var MagicValues = []uint32{
	0xD7B7_3E80,
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
