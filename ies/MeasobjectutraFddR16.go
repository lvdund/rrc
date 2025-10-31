package ies

// MeasObjectUTRA-FDD-r16 ::= SEQUENCE
// Extensible
type MeasobjectutraFddR16 struct {
	CarrierfreqR16         ArfcnValueutraFddR16
	UtraFddQOffsetrangeR16 *UtraFddQOffsetrangeR16
	CellstoremovelistR16   *UtraFddCellindexlistR16
	CellstoaddmodlistR16   *CellstoaddmodlistutraFddR16
}
