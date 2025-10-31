package ies

import "rrc/utils"

// CSI-IM-Resource-csi-IM-ResourceElementPattern-pattern0 ::= SEQUENCE
type CsiImResourceCsiImResourceelementpatternPattern0 struct {
	SubcarrierlocationP0 CsiImResourceCsiImResourceelementpatternPattern0SubcarrierlocationP0
	SymbollocationP0     utils.INTEGER `lb:0,ub:12`
}
