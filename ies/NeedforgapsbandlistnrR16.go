package ies

// NeedForGapsBandListNR-r16 ::= SEQUENCE OF NeedForGapsNR-r16
// SIZE (1..maxBands)
type NeedforgapsbandlistnrR16 struct {
	Value []NeedforgapsnrR16 `lb:1,ub:maxBands`
}
