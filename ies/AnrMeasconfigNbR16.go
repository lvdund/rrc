package ies

// ANR-MeasConfig-NB-r16 ::= SEQUENCE
// Extensible
type AnrMeasconfigNbR16 struct {
	AnrQualitythresholdR16 NrsrpRangeNbR14
	AnrCarrierlistR16      AnrCarrierlistNbR16
}
