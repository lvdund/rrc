package ies

// PSCellToAddMod-r12 ::= SEQUENCE
// Extensible
type PscelltoaddmodR12 struct {
	ScellindexR12                         ScellindexR10
	CellidentificationR12                 *PscelltoaddmodR12CellidentificationR12
	RadioresourceconfigcommonpscellR12    *RadioresourceconfigcommonpscellR12
	RadioresourceconfigdedicatedpscellR12 *RadioresourceconfigdedicatedpscellR12
}
