package ies

import "rrc/utils"

// InterFreqNeighCellInfo ::= SEQUENCE
// Extensible
type Interfreqneighcellinfo struct {
	Physcellid             Physcellid
	QOffsetcell            QOffsetrange
	QRxlevminoffsetcell    *utils.INTEGER `lb:0,ub:8`
	QRxlevminoffsetcellsul *utils.INTEGER `lb:0,ub:8`
	QQualminoffsetcell     *utils.INTEGER `lb:0,ub:8`
}
