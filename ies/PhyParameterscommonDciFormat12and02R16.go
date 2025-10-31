package ies

import "rrc/utils"

// Phy-ParametersCommon-dci-Format1-2And0-2-r16 ::= ENUMERATED
type PhyParameterscommonDciFormat12and02R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDciFormat12and02R16EnumeratedNothing = iota
	PhyParameterscommonDciFormat12and02R16EnumeratedSupported
)
