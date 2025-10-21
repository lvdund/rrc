package ies

import "rrc/utils"

// CellSelectionInfo-NB-v1430 ::= SEQUENCE
type CellselectioninfoNbV1430 struct {
	Powerclass14dbmOffsetR14 *utils.ENUMERATED
	CeAuthorisationoffsetR14 *utils.ENUMERATED
}
