package ies

import "rrc/utils"

// CE-Parameters-v1350-unicastFrequencyHopping-r13 ::= ENUMERATED
type CeParametersV1350UnicastfrequencyhoppingR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersV1350UnicastfrequencyhoppingR13EnumeratedNothing = iota
	CeParametersV1350UnicastfrequencyhoppingR13EnumeratedSupported
)
