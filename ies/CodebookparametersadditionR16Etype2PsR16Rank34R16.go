package ies

import "rrc/utils"

// CodebookParametersAddition-r16-etype2-PS-r16-rank3-4-r16 ::= ENUMERATED
type CodebookparametersadditionR16Etype2PsR16Rank34R16 struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersadditionR16Etype2PsR16Rank34R16EnumeratedNothing = iota
	CodebookparametersadditionR16Etype2PsR16Rank34R16EnumeratedSupported
)
