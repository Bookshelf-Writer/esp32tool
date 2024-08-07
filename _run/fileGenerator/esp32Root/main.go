package main

import (
	"github.com/Bookshelf-Writer/esp32tool/lib/generator"
	"io/ioutil"
)

//###########################################################//

const (
	ImportName  = "github.com/Bookshelf-Writer/esp32tool"
	PackageName = "root"

	FileMap = "esp32/root/map.go"
)

////

func main() {
	obj := generator.Init("", FileMap)

	//
	var folders []string
	files, _ := ioutil.ReadDir("esp32/" + PackageName + "/branch")

	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}

	//

	obj.PrintLN("import(")
	for _, folder := range folders {
		obj.Offset(1).String(ImportName + "/esp32/" + PackageName + "/branch/" + folder).LN()
	}
	obj.Offset(1).String(ImportName + "/esp32/root/patterns").LN()
	obj.Offset(1).String(ImportName + "/lib/output").LN()
	obj.Offset(1).String(ImportName + "/lib/serial").LN()
	obj.PrintLN(")").LN()

	//

	obj.PrintLN("var MagicMap map[uint32]func(*serial.SerialObj, *output.LogObj) (patterns.EspInterface, error)").LN()

	obj.PrintLN("func init() {")
	obj.Offset(1).PrintLN("MagicMap = make(map[uint32]func(*serial.SerialObj, *output.LogObj) (patterns.EspInterface, error))").LN()
	for _, folder := range folders {
		obj.Offset(1).PrintLN("for _, key := range " + folder + ".MagicValues {")
		obj.Offset(2).PrintLN("MagicMap[key] = " + folder + ".Init")
		obj.Offset(1).PrintLN("}")
	}
	obj.PrintLN("}").LN()

	//

	obj.Save(PackageName).End()
}
