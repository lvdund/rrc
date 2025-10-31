package ies

import "rrc/utils"

// Phy-ParametersCommon-csi-RS-CFRA-ForHO ::= ENUMERATED
type PhyParameterscommonCsiRsCfraForho struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCsiRsCfraForhoEnumeratedNothing = iota
	PhyParameterscommonCsiRsCfraForhoEnumeratedSupported
)
