package ies

// MeasTimingList ::= SEQUENCE OF MeasTiming
// SIZE (1..maxMeasFreqsMN)
type Meastiminglist struct {
	Value []Meastiming `lb:1,ub:maxMeasFreqsMN`
}
