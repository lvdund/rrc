package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-ce-MultiTB-FrequencyHopping-r16 ::= ENUMERATED
type CeMultitbParametersR16CeMultitbFrequencyhoppingR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16CeMultitbFrequencyhoppingR16EnumeratedNothing = iota
	CeMultitbParametersR16CeMultitbFrequencyhoppingR16EnumeratedSupported
)
