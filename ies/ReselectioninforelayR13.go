package ies

import "rrc/utils"

// ReselectionInfoRelay-r13 ::= SEQUENCE
type ReselectioninforelayR13 struct {
	QRxlevminR13         QRxlevmin
	FiltercoefficientR13 Filtercoefficient
	MinhystR13           *utils.ENUMERATED
}
