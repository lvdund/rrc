package ies

import "rrc/utils"

// Phy-ParametersCommon-configuredUL-GrantType1 ::= ENUMERATED
type PhyParameterscommonConfiguredulGranttype1 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonConfiguredulGranttype1EnumeratedNothing = iota
	PhyParameterscommonConfiguredulGranttype1EnumeratedSupported
)
