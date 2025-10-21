package ies

import "rrc/utils"

// IntraFreqCellReselectionInfo-NB-v1430 ::= SEQUENCE
type IntrafreqcellreselectioninfoNbV1430 struct {
	Powerclass14dbmOffsetR14 *utils.ENUMERATED
	CeAuthorisationoffsetR14 *utils.ENUMERATED
}
