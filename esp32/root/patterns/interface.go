package patterns

import (
	"github.com/Bookshelf-Writer/esp32tool/esp32/code"
)

//###########################################################//

type EspInterface interface {
	MagicValues() []uint32
	Name() string

	GetFrequency() code.FrequencyType
	GetSize() code.SizeType
}
