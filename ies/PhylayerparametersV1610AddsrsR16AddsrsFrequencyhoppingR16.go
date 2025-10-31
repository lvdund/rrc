package ies

import "rrc/utils"

// PhyLayerParameters-v1610-addSRS-r16-addSRS-FrequencyHopping-r16 ::= ENUMERATED
type PhylayerparametersV1610AddsrsR16AddsrsFrequencyhoppingR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610AddsrsR16AddsrsFrequencyhoppingR16EnumeratedNothing = iota
	PhylayerparametersV1610AddsrsR16AddsrsFrequencyhoppingR16EnumeratedSupported
)
