package ies

import "rrc/utils"

// PH-UplinkCarrierSCG-ph-Type1or3 ::= ENUMERATED
type PhUplinkcarrierscgPhType1or3 struct {
	Value utils.ENUMERATED
}

const (
	PhUplinkcarrierscgPhType1or3EnumeratedNothing = iota
	PhUplinkcarrierscgPhType1or3EnumeratedType1
	PhUplinkcarrierscgPhType1or3EnumeratedType3
)
