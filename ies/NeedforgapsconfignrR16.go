package ies

// NeedForGapsConfigNR-r16 ::= SEQUENCE
type NeedforgapsconfignrR16 struct {
	RequestedtargetbandfilternrR16 *[]Freqbandindicatornr `lb:1,ub:maxBands`
}
