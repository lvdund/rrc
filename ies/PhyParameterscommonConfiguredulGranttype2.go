package ies

import "rrc/utils"

// Phy-ParametersCommon-configuredUL-GrantType2 ::= ENUMERATED
type PhyParameterscommonConfiguredulGranttype2 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonConfiguredulGranttype2EnumeratedNothing = iota
	PhyParameterscommonConfiguredulGranttype2EnumeratedSupported
)
