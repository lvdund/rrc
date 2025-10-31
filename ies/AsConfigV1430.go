package ies

// AS-Config-v1430 ::= SEQUENCE
type AsConfigV1430 struct {
	SourceslV2xCommconfigR14 *SlV2xConfigdedicatedR14
	SourcelwaConfigR14       *LwaConfigR13
	SourcewlanMeasresultR14  *MeasresultlistwlanR13
}
