package ies

import "rrc/utils"

// Phy-ParametersCommon-pdsch-MappingTypeB ::= ENUMERATED
type PhyParameterscommonPdschMappingtypeb struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPdschMappingtypebEnumeratedNothing = iota
	PhyParameterscommonPdschMappingtypebEnumeratedSupported
)
