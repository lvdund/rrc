package ies

import "rrc/utils"

// Phy-ParametersCommon-type2-CG-ReleaseDCI-0-2-r16 ::= ENUMERATED
type PhyParameterscommonType2CgReleasedci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonType2CgReleasedci02R16EnumeratedNothing = iota
	PhyParameterscommonType2CgReleasedci02R16EnumeratedSupported
)
