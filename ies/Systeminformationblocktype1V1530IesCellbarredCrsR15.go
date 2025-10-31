package ies

import "rrc/utils"

// SystemInformationBlockType1-v1530-IEs-cellBarred-CRS-r15 ::= ENUMERATED
type Systeminformationblocktype1V1530IesCellbarredCrsR15 struct {
	Value utils.ENUMERATED
}

const (
	Systeminformationblocktype1V1530IesCellbarredCrsR15EnumeratedNothing = iota
	Systeminformationblocktype1V1530IesCellbarredCrsR15EnumeratedBarred
	Systeminformationblocktype1V1530IesCellbarredCrsR15EnumeratedNotbarred
)
