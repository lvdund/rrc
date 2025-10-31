package ies

// Cell-ToAddMod-r12 ::= SEQUENCE
// Extensible
type CellToaddmodR12 struct {
	ScellindexR12          ScellindexR10
	CellidentificationR12  *CellToaddmodR12CellidentificationR12
	MeasresultcelltoaddR12 *CellToaddmodR12MeasresultcelltoaddR12
}
