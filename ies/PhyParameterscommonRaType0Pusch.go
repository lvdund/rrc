package ies

import "rrc/utils"

// Phy-ParametersCommon-ra-Type0-PUSCH ::= ENUMERATED
type PhyParameterscommonRaType0Pusch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonRaType0PuschEnumeratedNothing = iota
	PhyParameterscommonRaType0PuschEnumeratedSupported
)
