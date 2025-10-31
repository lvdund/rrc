package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-FrequencyHopping-r16 ::= ENUMERATED
type PurParametersR16PurFrequencyhoppingR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurFrequencyhoppingR16EnumeratedNothing = iota
	PurParametersR16PurFrequencyhoppingR16EnumeratedSupported
)
