package esp32s3

import "github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"

//###########################################################//

const Name = "ESP32 S3"

var MagicValues = []uint32{
	0x9,
}

var (
	IROM = patterns.AddressObj{
		Offset: 0x4200_0000,
		Size:   0x4400_0000 - 0x4200_0000,
	}
	DROM = patterns.AddressObj{
		Offset: 0x3c00_0000,
		Size:   0x3e00_0000 - 0x3c00_0000,
	}
)

const Efuse uint32 = 0x6000_7000
