package code

//###########################################################//

type FrequencyType byte

const (
	Freq12MHz FrequencyType = iota + 1
	Freq15MHz
	Freq16MHz
	Freq20MHz
	Freq24MHz
	Freq26MHz
	Freq30MHz
	Freq40MHz
	Freq48MHz
	Freq60MHz
	Freq80MHz
	Freq160Mhz
	Freq240MHz
)

var FreqMap = map[FrequencyType]string{
	Freq12MHz:  "12MHz",
	Freq15MHz:  "15MHz",
	Freq16MHz:  "16MHz",
	Freq20MHz:  "20MHz",
	Freq24MHz:  "24MHz",
	Freq26MHz:  "26MHz",
	Freq30MHz:  "30MHz",
	Freq40MHz:  "40MHz",
	Freq48MHz:  "48MHz",
	Freq60MHz:  "60MHz",
	Freq80MHz:  "80MHz",
	Freq160Mhz: "160MHz",
	Freq240MHz: "240MHz",
}

func (f FrequencyType) String() string {
	str, ok := FreqMap[f]
	if ok {
		return str
	}
	return "unknown"
}
