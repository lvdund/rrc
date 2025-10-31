package ies

import "rrc/utils"

// CSI-IM-Resource-csi-IM-ResourceElementPattern-pattern1 ::= SEQUENCE
type CsiImResourceCsiImResourceelementpatternPattern1 struct {
	SubcarrierlocationP1 CsiImResourceCsiImResourceelementpatternPattern1SubcarrierlocationP1
	SymbollocationP1     utils.INTEGER `lb:0,ub:13`
}
