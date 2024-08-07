package esp32c3

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32 C3"

var MagicValues = []uint32{
	0x6921_506f, // ECO1 + ECO2
	0x1b31_506f, // ECO3
	0x4881_606F, // ECO6
	0x4361_606f, // ECO7
}

var (
	IROM = patterns.AddressObj{
		Offset: 0x4200_0000,
		Size:   0x4280_0000 - 0x4200_0000,
	}
	DROM = patterns.AddressObj{
		Offset: 0x3c00_0000,
		Size:   0x3c80_0000 - 0x3c00_0000,
	}
)

const Efuse uint32 = 0x6000_8800
