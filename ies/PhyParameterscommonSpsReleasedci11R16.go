package ies

import "rrc/utils"

// Phy-ParametersCommon-sps-ReleaseDCI-1-1-r16 ::= ENUMERATED
type PhyParameterscommonSpsReleasedci11R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSpsReleasedci11R16EnumeratedNothing = iota
	PhyParameterscommonSpsReleasedci11R16EnumeratedSupported
)
