package ies

import "rrc/utils"

// Phy-ParametersCommon-sps-ReleaseDCI-1-2-r16 ::= ENUMERATED
type PhyParameterscommonSpsReleasedci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSpsReleasedci12R16EnumeratedNothing = iota
	PhyParameterscommonSpsReleasedci12R16EnumeratedSupported
)
