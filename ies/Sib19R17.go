package ies

import "rrc/utils"

// SIB19-r17 ::= SEQUENCE
// Extensible
type Sib19R17 struct {
	NtnConfigR17                   *NtnConfigR17
	TServiceR17                    *utils.INTEGER `lb:0,ub:549755813887`
	ReferencelocationR17           *ReferencelocationR17
	DistancethreshR17              *utils.INTEGER `lb:0,ub:65525`
	NtnNeighcellconfiglistR17      *NtnNeighcellconfiglistR17
	Latenoncriticalextension       *utils.OCTETSTRING
	NtnNeighcellconfiglistextV1720 *NtnNeighcellconfiglistR17 `ext`
}
