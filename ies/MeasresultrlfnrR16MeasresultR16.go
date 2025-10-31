package ies

// MeasResultRLFNR-r16-measResult-r16 ::= SEQUENCE
type MeasresultrlfnrR16MeasresultR16 struct {
	CellresultsR16    *MeasresultrlfnrR16MeasresultR16CellresultsR16
	RsindexresultsR16 *MeasresultrlfnrR16MeasresultR16RsindexresultsR16 `lb:64,ub:64`
}
