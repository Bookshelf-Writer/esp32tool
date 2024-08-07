package esp32c2

import (
	"github.com/Bookshelf-Writer/esp32tool/esp32/root/patterns"
	"github.com/Bookshelf-Writer/esp32tool/lib/output"
	"github.com/Bookshelf-Writer/esp32tool/lib/serial"
)

//###########################################################//

type EspDevObj struct {
	patterns.EspInterface
}

func Init(*serial.SerialObj, *output.LogObj) (patterns.EspInterface, error) {
	obj := EspDevObj{}

	return &obj, nil
}

////

func (obj *EspDevObj) MagicValues() []uint32 {
	return MagicValues
}

func (obj *EspDevObj) Name() string {
	return Name
}
