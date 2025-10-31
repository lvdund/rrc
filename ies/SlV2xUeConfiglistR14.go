package ies

// SL-V2X-UE-ConfigList-r14 ::= SEQUENCE OF SL-V2X-InterFreqUE-Config-r14
// SIZE (1.. maxCellIntra)
type SlV2xUeConfiglistR14 struct {
	Value []SlV2xInterfrequeConfigR14 `lb:1,ub:maxCellIntra`
}
