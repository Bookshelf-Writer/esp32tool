package serial

import "errors"

//###########################################################//

const BaudRateMin = 100

var ErrBaudRateMin = errors.New("baudRate too small")
