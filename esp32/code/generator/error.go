package main

import "github.com/Bookshelf-Writer/esp32tool/lib/generator"

//###########################################################//

func errors() {
	obj := generator.Init("Err", "esp32/code/error.go")
	val := obj.NewVal()

	//

	val.Add(0x05, "Invalid message received")
	val.Add(0x06, "Bootloader failed to execute command")
	val.Add(0x07, "Received message has invalid CRC").Delim()

	val.Add(0x08, "Bootloader failed to write to flash")
	val.Add(0x09, "Bootloader failed to read from flash")
	val.Add(0x0A, "Invalid length for flash read")
	val.Add(0x0B, "Malformed compressed data received").Delim()

	val.Add(0xc0, "Bad data length")
	val.Add(0xc1, "Bad data checksum")
	val.Add(0xc2, "Bad block size")
	val.Add(0xc3, "Invalid command").Delim()

	val.Add(0xc4, "SPI operation failed")
	val.Add(0xc5, "SPI unlock failed").Delim()

	val.Add(0xc7, "Error when uncompressing the data")
	val.Add(0xc8, "Didn't receive enough data")
	val.Add(0xc9, "Received too much data")
	val.Add(0xff, "Other")

	//

	build(obj, val)

	//

	obj.LN().Print("func ErrorDetect(b byte) ").Name.Type().PrintLN(" {")
	obj.Offset(1).Print("_, ok := ").Name.Map().Print("[").Name.Type().PrintLN("(b)]")
	obj.Offset(1).PrintLN("if ok {").Offset(2).PrintLN("return ErrOther").Offset(1).PrintLN("}")
	obj.Offset(1).Print("return ").Name.Type().PrintLN("(b)")
	obj.PrintLN("}")

	//

	obj.Save("code").Add.Type(obj.Name.GetType(), "byte").End()
}
