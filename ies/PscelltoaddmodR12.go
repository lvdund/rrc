package ies

import "rrc/utils"

// PSCellToAddMod-r12 ::= SEQUENCE
// Extensible
type PscelltoaddmodR12 struct {
	ScellindexR12                         ScellindexR10
	CellidentificationR12                 *interface{}
	RadioresourceconfigcommonpscellR12    *RadioresourceconfigcommonpscellR12
	RadioresourceconfigdedicatedpscellR12 *RadioresourceconfigdedicatedpscellR12
}
