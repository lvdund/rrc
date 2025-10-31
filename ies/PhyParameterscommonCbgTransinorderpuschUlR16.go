package ies

import "rrc/utils"

// Phy-ParametersCommon-cbg-TransInOrderPUSCH-UL-r16 ::= ENUMERATED
type PhyParameterscommonCbgTransinorderpuschUlR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCbgTransinorderpuschUlR16EnumeratedNothing = iota
	PhyParameterscommonCbgTransinorderpuschUlR16EnumeratedSupported
)
