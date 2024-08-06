package root

import "github.com/Bookshelf-Writer/esp32tool/esp32/code"

//###########################################################//

type EspInterface interface {
	GetMagicValues() []uint32

	GetFrequency() code.FrequencyType
	GetSize() code.SizeType
	GetName() string
}
