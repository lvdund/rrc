package ies

import "rrc/utils"

// Phy-ParametersCommon-interleavingVRB-ToPRB-PDSCH ::= ENUMERATED
type PhyParameterscommonInterleavingvrbToprbPdsch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonInterleavingvrbToprbPdschEnumeratedNothing = iota
	PhyParameterscommonInterleavingvrbToprbPdschEnumeratedSupported
)
