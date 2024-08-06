package code

//###########################################################//

type FeatureType byte

const (
	FeatureWiFi FeatureType = iota + 1
	FeatureBT
	FeatureSingleCore
	FeatureDualCore
	FeatureEmbeddedFlash
	FeatureEmbeddedPSRAM
	FeatureVrefEfuse
	FeatureBLK3
	FeatureSchemeNone
	FeatureSchemeHP
	FeatureSchemeRepeat
	FeatureSchemeInvalid
)

var FeatureMap = map[FeatureType]string{
	FeatureWiFi:          "WiFi",
	FeatureBT:            "BT",
	FeatureSingleCore:    "Single Core",
	FeatureDualCore:      "Dual Core",
	FeatureEmbeddedFlash: "Embedded Flash",
	FeatureEmbeddedPSRAM: "Embedded PSRAM",
	FeatureVrefEfuse:     "VRef calibration in efuse",
	FeatureBLK3:          "BLK3 partially reserved",
	FeatureSchemeNone:    "Coding Scheme None",
	FeatureSchemeHP:      "Coding Scheme 3/4",
	FeatureSchemeRepeat:  "Coding Scheme Repeat (UNSUPPORTED)",
	FeatureSchemeInvalid: "Coding Scheme Invalid",
}

func (f FeatureType) String() string {
	str, ok := FeatureMap[f]
	if ok {
		return str
	}
	return "unknown"
}
