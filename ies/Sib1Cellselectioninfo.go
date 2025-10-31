package ies

import "rrc/utils"

// SIB1-cellSelectionInfo ::= SEQUENCE
type Sib1Cellselectioninfo struct {
	QRxlevmin       QRxlevmin
	QRxlevminoffset *utils.INTEGER `lb:0,ub:8`
	QRxlevminsul    *QRxlevmin
	QQualmin        *QQualmin
	QQualminoffset  *utils.INTEGER `lb:0,ub:8`
}
