package ies

// SCellToAddMod-r10 ::= SEQUENCE
// Extensible
type ScelltoaddmodR10 struct {
	ScellindexR10                        ScellindexR10
	CellidentificationR10                *ScelltoaddmodR10CellidentificationR10
	RadioresourceconfigcommonscellR10    *RadioresourceconfigcommonscellR10
	RadioresourceconfigdedicatedscellR10 *RadioresourceconfigdedicatedscellR10
}
