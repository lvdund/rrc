package ies

// AffectedCarrierFreqComb-r16 ::= SEQUENCE
type AffectedcarrierfreqcombR16 struct {
	AffectedcarrierfreqcombR16 *[]ArfcnValuenr `lb:2,ub:maxNrofServingCells`
	VictimsystemtypeR16        VictimsystemtypeR16
}
