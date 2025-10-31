package ies

// ReselectionInfoRelay-r13 ::= SEQUENCE
type ReselectioninforelayR13 struct {
	QRxlevminR13         QRxlevmin
	FiltercoefficientR13 Filtercoefficient
	MinhystR13           *ReselectioninforelayR13MinhystR13
}
