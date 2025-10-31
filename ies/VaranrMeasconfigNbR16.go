package ies

// VarANR-MeasConfig-NB-r16 ::= SEQUENCE
type VaranrMeasconfigNbR16 struct {
	AnrQualitythresholdR16 NrsrpRangeNbR14
	AnrCarrierlistR16      AnrCarrierlistNbR16
}
