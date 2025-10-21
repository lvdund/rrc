package ies

import "rrc/utils"

// BandParameters-r13 ::= SEQUENCE
type BandparametersR13 struct {
	BandeutraR13        FreqbandindicatorR11
	BandparametersulR13 *BandparametersulR13
	BandparametersdlR13 *BandparametersdlR13
	SupportedcsiProcR13 *utils.ENUMERATED
}
