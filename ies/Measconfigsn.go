package ies

// MeasConfigSN ::= SEQUENCE
// Extensible
type Measconfigsn struct {
	Measuredfrequenciessn *[]NrFreqinfo `lb:1,ub:maxMeasFreqsSN`
}
