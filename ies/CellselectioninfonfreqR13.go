package ies

import "rrc/utils"

// CellSelectionInfoNFreq-r13 ::= SEQUENCE
type CellselectioninfonfreqR13 struct {
	QRxlevminR13            QRxlevmin
	QRxlevminoffset         *utils.INTEGER
	QHystR13                utils.ENUMERATED
	QRxlevminreselectionR13 QRxlevmin
	TReselectioneutraR13    TReselection
}
