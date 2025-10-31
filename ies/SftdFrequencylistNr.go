package ies

// SFTD-FrequencyList-NR ::= SEQUENCE OF ARFCN-ValueNR
// SIZE (1..maxCellSFTD)
type SftdFrequencylistNr struct {
	Value []ArfcnValuenr `lb:1,ub:maxCellSFTD`
}
