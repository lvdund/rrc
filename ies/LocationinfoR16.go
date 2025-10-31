package ies

// LocationInfo-r16 ::= SEQUENCE
// Extensible
type LocationinfoR16 struct {
	CommonlocationinfoR16 *CommonlocationinfoR16
	BtLocationinfoR16     *LogmeasresultlistbtR16
	WlanLocationinfoR16   *LogmeasresultlistwlanR16
	SensorLocationinfoR16 *SensorLocationinfoR16
}
