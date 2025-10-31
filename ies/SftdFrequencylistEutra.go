package ies

// SFTD-FrequencyList-EUTRA ::= SEQUENCE OF ARFCN-ValueEUTRA
// SIZE (1..maxCellSFTD)
type SftdFrequencylistEutra struct {
	Value []ArfcnValueeutra `lb:1,ub:maxCellSFTD`
}
