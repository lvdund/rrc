package ies

// UplinkTxDirectCurrentMoreCarrierList-r17 ::= SEQUENCE OF CC-Group-r17
// SIZE (1..maxNrofCC-Group-r17)
type UplinktxdirectcurrentmorecarrierlistR17 struct {
	Value []CcGroupR17 `lb:1,ub:maxNrofCCGroupR17`
}
