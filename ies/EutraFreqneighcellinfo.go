package ies

import "rrc/utils"

// EUTRA-FreqNeighCellInfo ::= SEQUENCE
type EutraFreqneighcellinfo struct {
	Physcellid          EutraPhyscellid
	Dummy               EutraQOffsetrange
	QRxlevminoffsetcell *utils.INTEGER `lb:0,ub:8`
	QQualminoffsetcell  *utils.INTEGER `lb:0,ub:8`
}
