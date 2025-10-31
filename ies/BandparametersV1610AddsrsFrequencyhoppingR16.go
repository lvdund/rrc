package ies

import "rrc/utils"

// BandParameters-v1610-addSRS-FrequencyHopping-r16 ::= ENUMERATED
type BandparametersV1610AddsrsFrequencyhoppingR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1610AddsrsFrequencyhoppingR16EnumeratedNothing = iota
	BandparametersV1610AddsrsFrequencyhoppingR16EnumeratedSupported
)
