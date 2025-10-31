package ies

import "rrc/utils"

// Phy-ParametersCommon-type2-CG-ReleaseDCI-0-1-r16 ::= ENUMERATED
type PhyParameterscommonType2CgReleasedci01R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonType2CgReleasedci01R16EnumeratedNothing = iota
	PhyParameterscommonType2CgReleasedci01R16EnumeratedSupported
)
