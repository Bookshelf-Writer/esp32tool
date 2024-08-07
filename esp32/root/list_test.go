package root

import "testing"

//###########################################################//

func TestMap(t *testing.T) {
	for key, funcInit := range MagicMap {
		obj, _ := funcInit(nil, nil)
		t.Log(key, obj.Name(), obj.MagicValues())
	}
}
