package ies

import "rrc/utils"

// CodebookParametersAddition-r16-etype2-r16-amplitudeSubsetRestriction-r16 ::= ENUMERATED
type CodebookparametersadditionR16Etype2R16AmplitudesubsetrestrictionR16 struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersadditionR16Etype2R16AmplitudesubsetrestrictionR16EnumeratedNothing = iota
	CodebookparametersadditionR16Etype2R16AmplitudesubsetrestrictionR16EnumeratedSupported
)
