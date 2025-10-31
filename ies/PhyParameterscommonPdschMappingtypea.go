package ies

import "rrc/utils"

// Phy-ParametersCommon-pdsch-MappingTypeA ::= ENUMERATED
type PhyParameterscommonPdschMappingtypea struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPdschMappingtypeaEnumeratedNothing = iota
	PhyParameterscommonPdschMappingtypeaEnumeratedSupported
)
