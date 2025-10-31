package ies

// UplinkTxDirectCurrentList ::= SEQUENCE OF UplinkTxDirectCurrentCell
// SIZE (1..maxNrofServingCells)
type Uplinktxdirectcurrentlist struct {
	Value []Uplinktxdirectcurrentcell `lb:1,ub:maxNrofServingCells`
}
