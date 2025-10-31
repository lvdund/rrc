package ies

// UplinkTxDirectCurrentTwoCarrierList-r16 ::= SEQUENCE OF UplinkTxDirectCurrentTwoCarrier-r16
// SIZE (1..maxNrofTxDC-TwoCarrier-r16)
type UplinktxdirectcurrenttwocarrierlistR16 struct {
	Value []UplinktxdirectcurrenttwocarrierR16 `lb:1,ub:maxNrofTxDCTwocarrierR16`
}
