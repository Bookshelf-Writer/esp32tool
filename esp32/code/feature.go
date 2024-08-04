package code

/* This file is automatically generated */

type FeatureType byte

const (
	FeatureWiFi      FeatureType = 1 //0x01
	FeatureBluetooth FeatureType = 2 //0x02

	FeatureSingleCore FeatureType = 31 //0x1f
	FeatureDualCore   FeatureType = 32 //0x20

	FeatureClock80MHz  FeatureType = 51 //0x33
	FeatureClock160MHz FeatureType = 52 //0x34
	FeatureClock240MHz FeatureType = 53 //0x35
	FeatureClock320MHz FeatureType = 54 //0x36

	FeatureEmbeddedFlash         FeatureType = 101 //0x65
	FeatureVRefCalibration       FeatureType = 102 //0x66
	FeatureBLK3PartiallyReserved FeatureType = 103 //0x67

	FeatureCodingSchemeNone    FeatureType = 151 //0x97
	FeatureCodingScheme34      FeatureType = 152 //0x98
	FeatureCodingSchemeRepeat  FeatureType = 153 //0x99
	FeatureCodingSchemeInvalid FeatureType = 154 //0x9a
)

const (
	FeatureTextWiFi      = "WiFi"
	FeatureTextBluetooth = "Bluetooth"

	FeatureTextSingleCore = "Single Core"
	FeatureTextDualCore   = "Dual Core"

	FeatureTextClock80MHz  = "Clock 80MHz"
	FeatureTextClock160MHz = "Clock 160MHz"
	FeatureTextClock240MHz = "Clock 240MHz"
	FeatureTextClock320MHz = "Clock 320MHz"

	FeatureTextEmbeddedFlash         = "Embedded Flash"
	FeatureTextVRefCalibration       = "VRef calibration"
	FeatureTextBLK3PartiallyReserved = "BLK3 partially reserved"

	FeatureTextCodingSchemeNone    = "Coding Scheme None"
	FeatureTextCodingScheme34      = "Coding Scheme 3/4"
	FeatureTextCodingSchemeRepeat  = "Coding Scheme Repeat"
	FeatureTextCodingSchemeInvalid = "Coding Scheme Invalid"
)

var FeatureMap = map[FeatureType]string{
	FeatureWiFi:                  FeatureTextWiFi,
	FeatureBluetooth:             FeatureTextBluetooth,
	FeatureSingleCore:            FeatureTextSingleCore,
	FeatureDualCore:              FeatureTextDualCore,
	FeatureClock80MHz:            FeatureTextClock80MHz,
	FeatureClock160MHz:           FeatureTextClock160MHz,
	FeatureClock240MHz:           FeatureTextClock240MHz,
	FeatureClock320MHz:           FeatureTextClock320MHz,
	FeatureEmbeddedFlash:         FeatureTextEmbeddedFlash,
	FeatureVRefCalibration:       FeatureTextVRefCalibration,
	FeatureBLK3PartiallyReserved: FeatureTextBLK3PartiallyReserved,
	FeatureCodingSchemeNone:      FeatureTextCodingSchemeNone,
	FeatureCodingScheme34:        FeatureTextCodingScheme34,
	FeatureCodingSchemeRepeat:    FeatureTextCodingSchemeRepeat,
	FeatureCodingSchemeInvalid:   FeatureTextCodingSchemeInvalid,
}

func (obj FeatureType) String() string {
	val, ok := FeatureMap[obj]
	if ok {
		return val
	}
	return "Unknown FeatureType"
}