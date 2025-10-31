package ies

import "rrc/utils"

// CellSelectionInfoNFreq-r13 ::= SEQUENCE
type CellselectioninfonfreqR13 struct {
	QRxlevminR13            QRxlevmin
	QRxlevminoffset         *utils.INTEGER `lb:0,ub:8`
	QHystR13                CellselectioninfonfreqR13QHystR13
	QRxlevminreselectionR13 QRxlevmin
	TReselectioneutraR13    TReselection
}
