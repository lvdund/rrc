package ies

import "rrc/utils"

// SystemInformationBlockType1-cellSelectionInfo ::= SEQUENCE
type Systeminformationblocktype1Cellselectioninfo struct {
	QRxlevmin       QRxlevmin
	QRxlevminoffset *utils.INTEGER `lb:0,ub:8`
}
